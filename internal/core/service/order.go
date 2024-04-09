package service

import (
	"sync"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/repository"
)

type (
	OrderService struct {
		orderRepo         repository.IOrderRepository
		orderProductRepo  repository.IOrderProductRepository
		paymentResultRepo repository.IPaymentResultRepository
		productRepo       repository.IProductRepository
	}
)

func NewOrderService(repository repository.RepositoryCollection) port.OrderService {
	return &OrderService{
		orderRepo:         repository.OrderRepository,
		orderProductRepo:  repository.OrderProductRepository,
		paymentResultRepo: repository.PaymentResultRepository,
		productRepo:       repository.ProductReposotory,
	}
}

func (s OrderService) Create(form *domain.OrderDetail) (*domain.OrderDetail, *errs.AppError) {

	// validate stock
	for _, orderProduct := range form.OrderProducts {
		product, appErr := s.productRepo.GetOneByID(orderProduct.ProductID)
		if appErr != nil {
			return nil, appErr
		}

		if product.Stock < orderProduct.Quantity {
			logger.Error("Failed while create order: insufficient product stock")
			return nil, errs.NewBadRequestError("Insufficient product stock")
		}
	}

	form.CreatedAt = time.Now()
	form.UpdatedAt = time.Now()

	orderID, appErr := s.orderRepo.Insert(form)
	if appErr != nil {
		return nil, appErr
	}

	form.Order.OrderID = orderID
	return form, nil
}

func (s OrderService) GetList() ([]domain.OrderDetail, *errs.AppError) {
	orders, appErr := s.orderRepo.GetAll()
	if appErr != nil {
		return nil, appErr
	}

	return orders, nil
}

func (s OrderService) GetListByUser(userID int64) ([]domain.OrderDetail, *errs.AppError) {
	orders, appErr := s.orderRepo.GetAllByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	return orders, nil
}

func (s OrderService) GetDetail(ID int64) (*domain.OrderDetail, *errs.AppError) {

	var order *domain.OrderDetail
	var orderProducts []domain.OrderProduct

	orderChan := make(chan *domain.OrderDetail, 1)
	orderProductChan := make(chan []domain.OrderProduct, 1)
	errorChan := make(chan *errs.AppError, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		order, appErr := s.orderRepo.GetOneByID(ID)
		if appErr != nil {
			errorChan <- appErr
		}
		orderChan <- order
	}()

	go func() {
		defer wg.Done()
		orderProducts, appErr := s.orderProductRepo.GetAllByOrderID(ID)
		if appErr != nil {
			errorChan <- appErr
		}
		orderProductChan <- orderProducts
	}()

	wg.Wait()

	close(errorChan)
	close(orderChan)
	close(orderProductChan)

	for appErr := range errorChan {
		if appErr != nil {
			return nil, appErr
		}
	}

	for dataChan := range orderChan {
		order = dataChan
	}

	for dataChan := range orderProductChan {
		orderProducts = dataChan
	}

	order.OrderProducts = orderProducts
	return order, nil
}

func (s OrderService) UpdatePaid(form *domain.PaymentResult) *errs.AppError {

	_, appErr := s.orderRepo.GetOneByID(form.OrderID)
	if appErr != nil {
		return appErr
	}

	// check payment result by id
	checkPaymentResult, appErr := s.paymentResultRepo.CheckByID(form.PaymentResultID)
	if appErr != nil {
		return appErr
	}
	if checkPaymentResult {
		logger.Error("Failed while update order paid: payment result id already exist")
		return errs.NewBadRequestError("payment result id already exist")
	}

	// check order payment result
	checkPaymentResult, appErr = s.paymentResultRepo.CheckByOrderIDAndStatus(form.OrderID, "COMPLETED")
	if appErr != nil {
		return appErr
	}

	if checkPaymentResult {
		logger.Error("Failed while update order paid: order already paid")
		return errs.NewBadRequestError("order already paid")
	}

	appErr = s.orderRepo.UpdatePaid(form)
	if appErr != nil {
		return appErr
	}

	// update stock
	orderProducts, appErr := s.orderProductRepo.GetAllByOrderID(form.OrderID)
	if appErr != nil {
		return appErr
	}

	for _, orderProduct := range orderProducts {
		appErr := s.productRepo.UpdateStock(orderProduct.ProductID, orderProduct.Quantity)
		if appErr != nil {
			return appErr
		}

	}

	return nil
}

func (s OrderService) UpdateDelivered(ID int64) *errs.AppError {
	order, appErr := s.orderRepo.GetOneByID(ID)
	if appErr != nil {
		return appErr
	}

	if !order.IsPaid {
		logger.Error("Failed while update order delivered: order has not been paid")
		return errs.NewBadRequestError("order has not been paid")
	}

	if order.IsDelivered {
		logger.Error("Failed while update order delivered: order already deliverd")
		return errs.NewBadRequestError("order already deliverd")
	}

	appErr = s.orderRepo.UpdateDelivered(ID)
	if appErr != nil {
		return appErr
	}

	return nil
}

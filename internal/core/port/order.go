package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
)

type (
	OrderRepo interface {
		Insert(form *domain.OrderDetail) (int64, *errs.AppError)
		GetAll() ([]domain.OrderDetail, *errs.AppError)
		GetAllByUserID(userID int64) ([]domain.OrderDetail, *errs.AppError)
		GetOneByID(ID int64) (*domain.OrderDetail, *errs.AppError)
		UpdatePaid(form *domain.PaymentResult) *errs.AppError
		UpdateDelivered(ID int64) *errs.AppError
	}

	OrderService interface {
		Create(form *domain.OrderDetail) (*domain.OrderDetail, *errs.AppError)
		GetList() ([]domain.OrderDetail, *errs.AppError)
		GetListByUser(userID int64) ([]domain.OrderDetail, *errs.AppError)
		GetDetail(ID int64) (*domain.OrderDetail, *errs.AppError)
		UpdatePaid(form *domain.PaymentResult) *errs.AppError
		UpdateDelivered(ID int64) *errs.AppError
	}
)

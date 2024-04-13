package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/danisbagus/matchoshop/internal/usecase"
	"github.com/danisbagus/matchoshop/utils/auth"
	"github.com/danisbagus/matchoshop/utils/constants"
	"github.com/danisbagus/matchoshop/utils/helper"
	"github.com/gorilla/mux"
)

type OrderHandler struct {
	usecase usecase.IOrderUsecase
}

func (h OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	var req domain.CreateOrder

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := req.Validate()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	form := new(domain.OrderDetail)
	form.UserID = userInfo.UserID
	form.PaymentMethodID = req.PaymentMethodID
	form.ProductPrice = req.ProductPrice
	form.TaxPrice = req.TaxPrice
	form.ShippingPrice = req.ShippingPrice
	form.TotalPrice = req.TotalPrice
	form.ShipmentAddress.Address = req.ShippinmentAddress.Address
	form.ShipmentAddress.City = req.ShippinmentAddress.City
	form.ShipmentAddress.PostalCode = req.ShippinmentAddress.PostalCode
	form.ShipmentAddress.Country = req.ShippinmentAddress.Country

	orderProducts := make([]domain.OrderProduct, 0)
	for _, val := range req.OrderProduct {
		orderProduct := domain.OrderProduct{
			ProductID: val.ProductID,
			Quantity:  val.Quantity,
		}
		orderProducts = append(orderProducts, orderProduct)
	}

	form.OrderProducts = orderProducts

	createData, appErr := h.usecase.Create(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	resData := domain.NewOrderResponse(constants.SuccessCreate, createData)
	response.Write(w, http.StatusCreated, resData)
}

func (h OrderHandler) GetList(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	userID := userInfo.UserID

	orders, appErr := h.usecase.GetListByUser(userID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	resData := domain.NewGetOrderListResponse(constants.SuccesGet, orders)
	response.Write(w, http.StatusOK, resData)
}

func (h OrderHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	OrderID := helper.StringToInt64(vars["order_id"], 0)

	order, appErr := h.usecase.GetDetail(int64(OrderID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	resData := domain.NewGetOrderDetailResponse(constants.SuccesGet, order)
	response.Write(w, http.StatusOK, resData)
}

func (h OrderHandler) UpdatePaid(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	OrderID := helper.StringToInt64(vars["order_id"], 0)

	var req domain.UpdateOrderPaid

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := req.Validate()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	form := new(domain.PaymentResult)
	form.OrderID = OrderID
	form.PaymentResultID = req.PaymentMethodID
	form.Status = req.Status
	form.UpdateTime = helper.StringToDate(req.UpdateTime, time.RFC3339)
	form.Email = req.Email

	appErr = h.usecase.UpdatePaid(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	resData := domain.NewUpdatePaidResponse(constants.SuccessCreate, form)

	response.Write(w, http.StatusOK, resData)
}

func (h OrderHandler) UpdateDelivered(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := helper.StringToInt64(vars["order_id"], 0)

	appErr := h.usecase.UpdateDelivered(orderID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	resData := map[string]string{
		"message": constants.SuccessUpdate,
	}

	response.Write(w, http.StatusOK, resData)
}

func (h OrderHandler) GetListAdmin(w http.ResponseWriter, r *http.Request) {
	orders, appErr := h.usecase.GetList()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	resData := domain.NewGetOrderListResponse(constants.SuccesGet, orders)
	response.Write(w, http.StatusOK, resData)
}

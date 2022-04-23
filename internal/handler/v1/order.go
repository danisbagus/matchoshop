package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/utils/auth"
	"github.com/danisbagus/matchoshop/utils/constants"
	"github.com/danisbagus/matchoshop/utils/helper"
	"github.com/gorilla/mux"
)

type OrderHandler struct {
	Service port.OrderService
}

func (h OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	var req dto.CreateOrder

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

	createData, appErr := h.Service.Create(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	resData := dto.NewOrderResponse(constants.SuccessCreate, createData)
	response.Write(w, http.StatusCreated, resData)
}

func (h OrderHandler) GetList(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	userID := userInfo.UserID

	orders, appErr := h.Service.GetList(userID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	resData := dto.NewGetOrderListResponse(constants.SuccesGet, orders)
	response.Write(w, http.StatusOK, resData)
}

func (h OrderHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	OrderID := helper.StringToInt64(vars["order_id"], 0)

	order, appErr := h.Service.GetDetail(int64(OrderID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	resData := dto.NewGetOrderDetailResponse(constants.SuccesGet, order)
	response.Write(w, http.StatusOK, resData)
}

func (h OrderHandler) UpdatePaid(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	OrderID := helper.StringToInt64(vars["order_id"], 0)

	var req dto.UpdateOrderPaid

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

	appErr = h.Service.UpdatePaid(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	resData := dto.NewUpdatePaidResponse(constants.SuccessCreate, form)

	response.Write(w, http.StatusOK, resData)

}

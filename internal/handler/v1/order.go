package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/utils/auth"
	"github.com/danisbagus/matchoshop/utils/constants"
	"github.com/danisbagus/matchoshop/utils/helper"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	service port.OrderService
}

func NewOrderHandler(sevice port.OrderService) *OrderHandler {
	return &OrderHandler{service: sevice}
}

func (h OrderHandler) Create(c echo.Context) error {
	userInfo := auth.GetClaimData(c)
	var req dto.CreateOrder
	err := c.Bind(&req)
	if err != nil {
		logger.Error("Error while decoding create order request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	appErr := req.Validate()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
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

	createData, appErr := h.service.Create(form)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	resData := dto.NewOrderResponse(constants.SuccessCreate, createData)
	return c.JSON(http.StatusOK, resData)
}

func (h OrderHandler) GetList(c echo.Context) error {
	userInfo := auth.GetClaimData(c)
	userID := userInfo.UserID

	orders, appErr := h.service.GetListByUser(userID)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	resData := dto.NewGetOrderListResponse(constants.SuccesGet, orders)
	return c.JSON(http.StatusOK, resData)
}

func (h OrderHandler) GetDetail(c echo.Context) error {
	orderID, _ := strconv.Atoi(c.Param("order_id"))

	order, appErr := h.service.GetDetail(int64(orderID))
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	resData := dto.NewGetOrderDetailResponse(constants.SuccesGet, order)
	return c.JSON(http.StatusOK, resData)
}

func (h OrderHandler) UpdatePaid(c echo.Context) error {
	orderID := helper.StringToInt64(c.Param("order_id"), 0)
	var req dto.UpdateOrderPaid

	err := c.Bind(&req)
	if err != nil {
		logger.Error("Error while decoding update order paid request: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	appErr := req.Validate()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	form := new(domain.PaymentResult)
	form.OrderID = orderID
	form.PaymentResultID = req.PaymentMethodID
	form.Status = req.Status
	form.UpdateTime = helper.StringToDate(req.UpdateTime, time.RFC3339)
	form.Email = req.Email

	appErr = h.service.UpdatePaid(form)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	resData := dto.NewUpdatePaidResponse(constants.SuccessCreate, form)

	return c.JSON(http.StatusOK, resData)
}

func (h OrderHandler) UpdateDelivered(c echo.Context) error {
	orderID := helper.StringToInt64(c.Param("order_id"), 0)

	appErr := h.service.UpdateDelivered(orderID)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	resData := map[string]string{
		"message": constants.SuccessUpdate,
	}

	return c.JSON(http.StatusOK, resData)
}

func (h OrderHandler) GetListAdmin(c echo.Context) error {
	orders, appErr := h.service.GetList()
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	resData := dto.NewGetOrderListResponse(constants.SuccesGet, orders)
	return c.JSON(http.StatusOK, resData)
}

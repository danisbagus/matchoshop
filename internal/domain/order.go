package domain

import (
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/domain/common/constants"
	"github.com/danisbagus/matchoshop/utils/helper"
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	OrderModel struct {
		OrderID         int64      `db:"order_id"`
		UserID          int64      `db:"user_id"`
		PaymentMethodID int64      `db:"payment_method_id"`
		ProductPrice    int64      `db:"product_price"`
		TaxPrice        int64      `db:"tax_price"`
		ShippingPrice   int64      `db:"shipping_price"`
		TotalPrice      int64      `db:"total_price"`
		IsPaid          bool       `db:"is_paid"`
		PaidAt          *time.Time `db:"paid_at"`
		IsDelivered     bool       `db:"is_delivered"`
		DeliveredAt     *time.Time `db:"delivered_at"`
		CreatedAt       time.Time  `db:"created_at"`
		UpdatedAt       time.Time  `db:"updated_at"`
	}

	OrderProductModel struct {
		OrderID   int64 `db:"order_id"`
		ProductID int64 `db:"product_id"`
		Quantity  int64 `db:"quantity"`
	}

	OrderDetail struct {
		OrderModel
		PaymentMethodName string `db:"payment_method_name"`
		UserName          string `db:"user_name"`
		UserEmail         string `db:"user_email"`
		ShipmentAddress
		OrderProducts []OrderProduct
		PaymentResult
	}

	CreateOrder struct {
		UserID             int64           `json:"-"`
		PaymentMethodID    int64           `json:"payment_method_id"`
		ProductPrice       int64           `json:"product_price"`
		TaxPrice           int64           `json:"tax_price"`
		ShippingPrice      int64           `json:"shipping_price"`
		TotalPrice         int64           `json:"total_price"`
		ShippinmentAddress ShipmentAddress `json:"shipment_address"`
		OrderProduct       []OrderProduct  `json:"order_product"`
	}

	OrderProduct struct {
		OrderID   int64  // need adjust
		ProductID int64  `json:"product_id"`
		Name      string `json:"name"`
		Image     string `json:"image"`
		Price     int64  `json:"price"`
		Quantity  int64  `json:"quantity"`
	}

	ShipmentAddress struct {
		Address    string `json:"address"`
		City       string `json:"city"`
		PostalCode string `json:"postal_code"`
		Country    string `json:"country"`
	}

	UpdateOrderPaid struct {
		OrderID         int64  `json:"-"`
		PaymentMethodID string `json:"payment_method_id"`
		Status          string `json:"status"`
		UpdateTime      string `json:"update_time"`
		Email           string `json:"email"`
	}

	CreateOrderResponse struct {
		OrderID int64 `json:"order_id"`
	}

	OrderDetailResponse struct {
		OrderID            int64           `json:"order_id"`
		PaymentMethodID    int64           `json:"payment_method_id"`
		PaymentMethodName  string          `json:"payment_method_name"`
		ProductPrice       int64           `json:"product_price"`
		TaxPrice           int64           `json:"tax_price"`
		ShippingPrice      int64           `json:"shipping_price"`
		TotalPrice         int64           `json:"total_price"`
		UserName           string          `json:"user_name"`
		UserEmail          string          `json:"user_email"`
		IsPaid             bool            `json:"is_paid"`
		PaidAt             string          `json:"paid_at"`
		IsDelivered        bool            `json:"is_delivered"`
		DeliveredAt        string          `json:"deliverd_at"`
		ShippinmentAddress ShipmentAddress `json:"shipment_address"`
		OrderProduct       []OrderProduct  `json:"order_product"`
	}

	OrderListResponse struct {
		OrderID         int64  `json:"order_id"`
		PaymentMethodID int64  `json:"payment_method_id"`
		ProductPrice    int64  `json:"product_price"`
		TaxPrice        int64  `json:"tax_price"`
		ShippingPrice   int64  `json:"shipping_price"`
		TotalPrice      int64  `json:"total_price"`
		UserName        string `json:"user_name"`
		IsPaid          bool   `json:"is_paid"`
		PaidAt          string `json:"paid_at"`
		IsDelivered     bool   `json:"is_delivered"`
		DeliveredAt     string `json:"deliverd_at"`
		CreatedAt       string `json:"created_at"`
	}

	UpdateOrderPaidResponse struct {
		OrderID         int64  `json:"order_id"`
		PaymentResultID string `json:"payment_result_id"`
	}
)

func NewOrderResponse(message string, data *OrderDetail) *ResponseData {
	resData := new(CreateOrderResponse)
	resData.OrderID = data.OrderModel.OrderID

	return GenerateResponseData(message, resData)
}

func NewGetOrderListResponse(message string, data []OrderDetail) *ResponseData {
	resData := make([]OrderListResponse, 0)

	for _, orderDetail := range data {
		var resOrderList OrderListResponse
		resOrderList.OrderID = orderDetail.OrderModel.OrderID
		resOrderList.PaymentMethodID = orderDetail.PaymentMethodID
		resOrderList.ProductPrice = orderDetail.ProductPrice
		resOrderList.TaxPrice = orderDetail.TaxPrice
		resOrderList.ShippingPrice = orderDetail.ShippingPrice
		resOrderList.TotalPrice = orderDetail.TotalPrice
		resOrderList.UserName = orderDetail.UserName
		resOrderList.IsPaid = orderDetail.IsPaid
		resOrderList.PaidAt = helper.PointDateToString(orderDetail.PaidAt, constants.DATE_FORMAT)
		resOrderList.IsDelivered = orderDetail.IsDelivered
		resOrderList.DeliveredAt = helper.PointDateToString(orderDetail.DeliveredAt, constants.DATE_FORMAT)
		resOrderList.CreatedAt = helper.PointDateToString(&orderDetail.CreatedAt, constants.DATE_FORMAT)
		resData = append(resData, resOrderList)
	}
	return GenerateResponseData(message, resData)
}

func NewGetOrderDetailResponse(message string, data *OrderDetail) *ResponseData {
	resData := new(OrderDetailResponse)
	resData.OrderID = data.OrderModel.OrderID
	resData.PaymentMethodID = data.PaymentMethodID
	resData.PaymentMethodName = data.PaymentMethodName
	resData.ProductPrice = data.ProductPrice
	resData.TaxPrice = data.TaxPrice
	resData.ShippingPrice = data.ShippingPrice
	resData.TotalPrice = data.TotalPrice
	resData.UserName = data.UserName
	resData.UserEmail = data.UserEmail
	resData.IsPaid = data.IsPaid
	resData.PaidAt = helper.PointDateToString(data.PaidAt, constants.DATE_FORMAT)
	resData.IsDelivered = data.IsDelivered
	resData.DeliveredAt = helper.PointDateToString(data.DeliveredAt, constants.DATE_FORMAT)
	resData.ShippinmentAddress = ShipmentAddress{
		Address:    data.ShipmentAddress.Address,
		City:       data.ShipmentAddress.City,
		PostalCode: data.ShipmentAddress.PostalCode,
		Country:    data.ShipmentAddress.Country,
	}

	orderProducts := make([]OrderProduct, 0)
	for _, value := range data.OrderProducts {
		var orderProduct OrderProduct
		orderProduct.ProductID = value.ProductID
		orderProduct.Price = value.Price
		orderProduct.Name = value.Name
		orderProduct.Image = value.Image
		orderProduct.Quantity = value.Quantity

		orderProducts = append(orderProducts, orderProduct)
	}
	resData.OrderProduct = orderProducts
	return GenerateResponseData(message, resData)
}

func NewUpdatePaidResponse(message string, data *PaymentResult) *ResponseData {
	resData := new(UpdateOrderPaidResponse)
	resData.OrderID = data.OrderID
	resData.PaymentResultID = data.PaymentResultID

	return GenerateResponseData(message, resData)
}

func (r CreateOrder) Validate() *errs.AppError {
	if err := validation.Validate(r.PaymentMethodID, validation.Required); err != nil {
		return errs.NewBadRequestError("payment method id is required")
	} else if err := validation.Validate(r.ProductPrice, validation.Required); err != nil {
		return errs.NewBadRequestError("product price is required")
	} else if err := validation.Validate(r.ProductPrice, validation.Min(100)); err != nil {
		return errs.NewBadRequestError("product price must more than equal 100")
	} else if err := validation.Validate(r.TaxPrice, validation.Min(0)); err != nil {
		return errs.NewBadRequestError("tax price must more than equal 0")
	} else if err := validation.Validate(r.ShippingPrice, validation.Min(0)); err != nil {
		return errs.NewBadRequestError("shipping price must more than equal 0")
	} else if err := validation.Validate(r.TotalPrice, validation.Required); err != nil {
		return errs.NewBadRequestError("total price is required")
	} else if err := validation.Validate(r.TotalPrice, validation.Min(100)); err != nil {
		return errs.NewBadRequestError("total price must more than equal 100")
	} else if err := validation.Validate(r.ShippinmentAddress, validation.Required); err != nil {
		return errs.NewBadRequestError("shipping address is required")
	} else if err := validation.Validate(r.ShippinmentAddress.Address, validation.Required); err != nil {
		return errs.NewBadRequestError("address is required")
	} else if err := validation.Validate(r.ShippinmentAddress.City, validation.Required); err != nil {
		return errs.NewBadRequestError("city address is required")
	} else if err := validation.Validate(r.ShippinmentAddress.PostalCode, validation.Required); err != nil {
		return errs.NewBadRequestError("postal code is required is required")
	} else if err := validation.Validate(r.ShippinmentAddress.Country, validation.Required); err != nil {
		return errs.NewBadRequestError("country is required is required")
	} else if err := validation.Validate(r.OrderProduct, validation.Required); err != nil {
		return errs.NewBadRequestError("order product is required")
	} else if err := validation.Validate(r.OrderProduct, validation.Required); err != nil || len(r.OrderProduct) <= 0 {
		return errs.NewBadRequestError("order product is required")
	}
	return nil
}

func (r UpdateOrderPaid) Validate() *errs.AppError {
	if err := validation.Validate(r.Status, validation.Required); err != nil {
		return errs.NewBadRequestError("status is required")
	} else if err := validation.Validate(r.UpdateTime, validation.Required); err != nil {
		return errs.NewBadRequestError("update time is required")
	}
	return nil
}

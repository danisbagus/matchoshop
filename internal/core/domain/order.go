package domain

import "time"

type (
	Order struct {
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

	OrderDetail struct {
		Order
		PaymentMethodName string `db:"payment_method_name"`
		UserName          string `db:"user_name"`
		UserEmail         string `db:"user_email"`
		ShipmentAddress
		OrderProducts []OrderProduct
		PaymentResult
	}
)

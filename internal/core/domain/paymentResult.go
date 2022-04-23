package domain

import "time"

type PaymentResult struct {
	PaymentResultID string    `db:"payment_result_id"`
	OrderID         int64     `db:"order_id"`
	Status          string    `db:"status"`
	UpdateTime      time.Time `db:"update_time"`
	Email           string    `db:"email"`
}

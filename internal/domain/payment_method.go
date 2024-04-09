package domain

type (
	PaymentMethod struct {
		PaymentMethodID int64  `db:"payment_result_id"`
		Name            string `db:"name"`
	}
)

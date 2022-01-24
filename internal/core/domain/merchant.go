package domain

type Merchant struct {
	MerchantID int64  `db:"merchant_id"`
	UserID     int64  `db:"user_id"`
	Name       string `db:"name"`
	Identifier string `db:"identifier"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
}

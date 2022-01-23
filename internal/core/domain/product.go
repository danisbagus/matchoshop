package domain

type Product struct {
	ProductID   int64  `db:"product_id"`
	Name        string `db:"name"`
	Sku         string `db:"sku"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

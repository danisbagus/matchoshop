package domain

type Product struct {
	ProductID   int64  `db:"product_id"`
	MerchantID  int64  `db:"merchant_id"`
	Name        string `db:"name"`
	Sku         string `db:"sku"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

type ProductList struct {
	ProductID           int64  `db:"product_id"`
	MerchantID          int64  `db:"merchant_id"`
	Name                string `db:"name"`
	Sku                 string `db:"sku"`
	Price               int64  `db:"price"`
	ProductCategoryName string `db:"product_category_name"`
}

type ProductDetail struct {
	ProductID         int64  `db:"product_id"`
	MerchantID        int64  `db:"merchant_id"`
	Name              string `db:"name"`
	Sku               string `db:"sku"`
	Price             int64  `db:"price"`
	Description       string `db:"description"`
	ProductCategories []ProductCategory
}

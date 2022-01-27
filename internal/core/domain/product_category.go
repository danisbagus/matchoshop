package domain

type ProductCategory struct {
	ProductCategoryID int64  `db:"product_category_id"`
	MerchantID        int64  `db:"merchant_id"`
	Name              string `db:"name"`
	CreatedAt         string `db:"created_at"`
	UpdatedAt         string `db:"updated_at"`
}

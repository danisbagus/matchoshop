package domain

type ProductProductCategory struct {
	ProductID         int64 `db:"product_id"`
	ProductCategoryID int64 `db:"product_category_id"`
}

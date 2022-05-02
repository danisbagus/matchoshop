package domain

type ProductModel struct {
	ProductID   int64
	Name        string
	Sku         string
	Brand       *string
	Image       *string
	Description *string
	Price       int64
	Stock       int64
	CreatedAt   string
	UpdatedAt   string
}

type Product struct {
	ProductModel
	ProductCategoryIDs []int64
}

type ProductList struct {
	ProductModel
	ProductCategoryID   int64
	ProductCategoryName string
}

type ProductDetail struct {
	ProductModel
	ProductCategories []ProductCategory
}

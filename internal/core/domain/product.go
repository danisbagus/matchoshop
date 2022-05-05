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
	Rating              float32
	NumbReviews         int64
}

type ProductDetail struct {
	ProductModel
	Rating            float32
	NumbReviews       int64
	ProductCategories []ProductCategory
	Review            []Review
}

type ProductListCriteria struct {
	Keyword string
	Page    int64
	Limit   int64
	Sort    string
	Order   string
}

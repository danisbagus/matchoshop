package domain

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/utils/helper"
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	ProductModel struct {
		ProductID   int64   `db:"product_id"`
		Name        string  `db:"name"`
		Sku         string  `db:"sku"`
		Brand       *string `db:"brand"`
		Image       *string `db:"image"`
		Description *string `db:"description"`
		Price       int64   `db:"price"`
		Stock       int64   `db:"stock"`
		CreatedAt   string  `db:"created_at"`
		UpdatedAt   string  `db:"updated_at"`
	}

	ProductProductCategoryModel struct {
		ProductID         int64 `db:"product_id"`
		ProductCategoryID int64 `db:"product_category_id"`
	}

	Product struct {
		ProductModel
		ProductCategoryIDs []int64
	}

	ProductList struct {
		ProductModel
		ProductCategoryID   int64
		ProductCategoryName string
		Rating              float32
		NumbReviews         int64
	}

	ProductDetail struct {
		ProductModel
		Rating            float32
		NumbReviews       int64
		ProductCategories []ProductCategoryModel
		Review            []Review
	}

	ProductListCriteria struct {
		Keyword string
		Page    int64
		Limit   int64
		Sort    string
		Order   string
	}

	ProductRequest struct {
		Name               string  `json:"name"`
		Sku                string  `json:"sku"`
		Brand              *string `json:"brand"`
		Image              *string `json:"Image"`
		Description        *string `json:"description"`
		ProductCategoryIDs []int64 `json:"product_category_id"`
		Price              int64   `json:"price"`
		Stock              int64   `json:"stock"`
	}

	ProductResponse struct {
		ProductID int64   `json:"product_id"`
		Name      string  `json:"name"`
		Sku       string  `json:"sku"`
		Brand     *string `json:"brand"`
		Image     *string `json:"image"`
		Price     int64   `json:"price"`
	}

	ProductListResponse struct {
		ProductResponse
		Rating            float32                   `json:"rating"`
		NumbReviews       int64                     `json:"numb_reviews"`
		ProductCategories []ProductCategoryResponse `json:"product_categories"`
	}

	ProductDetailtResponse struct {
		ProductResponse
		Description       *string                   `json:"description"`
		Stock             int64                     `json:"stock"`
		Rating            float32                   `json:"rating"`
		NumbReviews       int64                     `json:"numb_reviews"`
		ProductCategories []ProductCategoryResponse `json:"product_categories"`
		Review            []ReviewResponse          `json:"reviews"`
	}

	ResponsePaginateData struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}
)

func GenerateResponsePaginateData(message string, data interface{}, meta interface{}) *ResponsePaginateData {
	return &ResponsePaginateData{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func NewGetProductListResponse(message string, data []ProductDetail, meta *helper.Meta) *ResponsePaginateData {
	products := make([]ProductListResponse, 0)
	for _, value := range data {
		var product ProductListResponse
		product.ProductID = value.ProductID
		product.Name = value.Name
		product.Sku = value.Sku
		product.Image = value.Image
		product.Brand = value.Brand
		product.Price = value.Price
		product.Rating = value.Rating
		product.NumbReviews = value.NumbReviews

		productCategories := make([]ProductCategoryResponse, 0)
		for _, valueCategory := range value.ProductCategories {
			var category ProductCategoryResponse
			category.ProductCategoryID = valueCategory.ProductCategoryID
			category.Name = valueCategory.Name
			productCategories = append(productCategories, category)
		}

		product.ProductCategories = productCategories
		products = append(products, product)
	}
	return GenerateResponsePaginateData(message, products, meta)
}

func NewGetProductDetailResponse(message string, data *ProductDetail) *ResponseData {
	product := new(ProductDetailtResponse)
	product.ProductID = data.ProductID
	product.Name = data.Name
	product.Sku = data.Sku
	product.Image = data.Image
	product.Description = data.Description
	product.Brand = data.Brand
	product.Price = data.Price
	product.Rating = data.Rating
	product.NumbReviews = data.NumbReviews
	product.Stock = data.Stock

	productCategories := make([]ProductCategoryResponse, 0)
	for _, valData := range data.ProductCategories {
		productCategory := ProductCategoryResponse{
			ProductCategoryID: valData.ProductCategoryID,
			Name:              valData.Name,
		}
		productCategories = append(productCategories, productCategory)
	}
	product.ProductCategories = productCategories

	productReviews := make([]ReviewResponse, 0)
	for _, valData := range data.Review {
		productReview := ReviewResponse{
			ReviewID:  valData.ReviewID,
			UserID:    valData.UserID,
			UserName:  valData.UserName,
			ProductID: valData.ProductID,
			Rating:    valData.Rating,
			Comment:   valData.Comment,
			CreatedAt: valData.CreatedAt,
		}
		productReviews = append(productReviews, productReview)
	}

	product.Review = productReviews

	return GenerateResponseData(message, product)
}

func (r ProductRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Name, validation.Required); err != nil {
		return errs.NewBadRequestError("Name is required")
	} else if err := validation.Validate(r.Sku, validation.Required); err != nil {
		return errs.NewBadRequestError("SKU is required")
	} else if err := validation.Validate(r.Brand, validation.Required); err != nil {
		return errs.NewBadRequestError("Brand is required")
	} else if err := validation.Validate(r.ProductCategoryIDs, validation.Required); err != nil {
		return errs.NewBadRequestError("Product category ID is required")
	} else if err := validation.Validate(r.Price, validation.Required); err != nil {
		return errs.NewBadRequestError("Price is required")
	} else if r.Price < 100 {
		return errs.NewValidationError("Minimum price is 100")
	} else if len(r.ProductCategoryIDs) < 1 {
		return errs.NewValidationError("Product category ID required")
	}
	return nil
}

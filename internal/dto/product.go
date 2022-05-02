package dto

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ProductRequest struct {
	Name               string  `json:"name"`
	Sku                string  `json:"sku"`
	Brand              *string `json:"brand"`
	Image              *string `json:"Image"`
	Description        *string `json:"description"`
	ProductCategoryIDs []int64 `json:"product_category_id"`
	Price              int64   `json:"price"`
	Stock              int64   `json:"stock"`
}

type ProductResponse struct {
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Sku       string  `json:"sku"`
	Brand     *string `json:"brand"`
	Image     *string `json:"image"`
	Price     int64   `json:"price"`
}

type ProductListResponse struct {
	ProductResponse
	Rating            float32                   `json:"rating"`
	NumbReviews       int64                     `json:"numb_reviews"`
	ProductCategories []ProductCategoryResponse `json:"product_categories"`
}

type ProductDetailtResponse struct {
	ProductResponse
	Description       *string                   `json:"description"`
	Stock             int64                     `json:"stock"`
	Rating            float32                   `json:"rating"`
	NumbReviews       int64                     `json:"numb_reviews"`
	ProductCategories []ProductCategoryResponse `json:"product_categories"`
}

func NewGetProductListResponse(message string, data []domain.ProductDetail) *ResponseData {
	products := make([]ProductListResponse, 0)
	for _, value := range data {
		var product ProductListResponse
		product.ProductID = value.ProductID
		product.Name = value.Name
		product.Sku = value.Sku
		product.Image = value.Image
		product.Brand = value.Brand
		product.Price = value.Price
		product.Rating = 4.2
		product.NumbReviews = 175

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
	return GenerateResponseData(message, products)
}

func NewGetProductDetailResponse(message string, data *domain.ProductDetail) *ResponseData {
	product := new(ProductDetailtResponse)
	product.ProductID = data.ProductID
	product.Name = data.Name
	product.Sku = data.Sku
	product.Image = data.Image
	product.Description = data.Description
	product.Brand = data.Brand
	product.Price = data.Price
	product.Rating = 4.2
	product.NumbReviews = 175
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

package dto

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateProductRequest struct {
	Name              string  `json:"name"`
	Sku               string  `json:"sku"`
	Description       string  `json:"description"`
	ProductCategoryID []int64 `json:"product_category_id"`
	Price             int64   `json:"price"`
}

type CreateProductResponse struct {
	ProductID int64 `json:"product_id"`
}

type ProductListResponse struct {
	ProductID         int64    `json:"product_id"`
	Name              string   `json:"name"`
	Sku               string   `json:"sku"`
	Image             string   `json:"image"`
	Price             int64    `json:"price"`
	Rating            float32  `json:"rating"`
	NumbReviews       int64    `json:"numb_reviews"`
	ProductCategories []string `json:"product_categories"`
}

type ProductDetailtResponse struct {
	ProductID         int64                     `json:"product_id"`
	Name              string                    `json:"name"`
	Sku               string                    `json:"sku"`
	Image             string                    `json:"image"`
	Price             int64                     `json:"price"`
	Description       string                    `json:"description"`
	Brand             string                    `json:"brand"`
	Stock             int64                     `json:"stock"`
	Rating            float32                   `json:"rating"`
	NumbReviews       int64                     `json:"numb_reviews"`
	ProductCategories []ProductCategoryResponse `json:"product_categories"`
}

func NewCreateProductResponse(message string, data *domain.Product) *ResponseData {

	createProductResponse := &CreateProductResponse{
		ProductID: data.ProductID,
	}

	return GenerateResponseData(message, createProductResponse)
}

func NewGetProductListResponse(message string, data []domain.ProductList) *ResponseData {

	products := make([]ProductListResponse, 0)
	mapProduct := make(map[int64]ProductListResponse)

	for i := 0; i < len(data); i++ {
		if mapValue, ok := mapProduct[data[i].ProductID]; ok {
			mapValue.ProductCategories = append(mapValue.ProductCategories, data[i].ProductCategoryName)
			mapProduct[data[i].ProductID] = mapValue
		} else {
			product := ProductListResponse{
				ProductID:         data[i].ProductID,
				Name:              data[i].Name,
				Sku:               data[i].Sku,
				Image:             "https://i.picsum.photos/id/167/300/300.jpg?hmac=pUvtrq_DnPmk3ui2smKHRtf7_EqebkD3wjh_9pE22d0",
				Price:             data[i].Price,
				Rating:            4.2,
				NumbReviews:       175,
				ProductCategories: []string{data[i].ProductCategoryName},
			}
			mapProduct[data[i].ProductID] = product
		}
	}

	for _, valData := range mapProduct {
		products = append(products, valData)
	}

	return GenerateResponseData(message, products)
}

func NewGetProductDetailResponse(message string, data *domain.ProductDetail) *ResponseData {
	productDetail := &ProductDetailtResponse{
		ProductID:   data.ProductID,
		Name:        data.Name,
		Sku:         data.Sku,
		Image:       "https://i.picsum.photos/id/399/1000/1000.jpg?hmac=Ily1BaUSN3DBCaX_fHgQkjQhqzeRhpY4zjKhwOYuA2E",
		Price:       data.Price,
		Description: data.Description,
		Brand:       "Other",
		Rating:      4.2,
		NumbReviews: 175,
		Stock:       44,
	}

	productCategories := make([]ProductCategoryResponse, 0)

	for _, valData := range data.ProductCategories {
		productCategory := ProductCategoryResponse{
			ProductCategoryID: valData.ProductCategoryID,
			Name:              valData.Name,
		}
		productCategories = append(productCategories, productCategory)
	}

	productDetail.ProductCategories = productCategories

	return GenerateResponseData(message, productDetail)
}

func (r CreateProductRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Name, validation.Required); err != nil {
		return errs.NewBadRequestError("Name is required")
	} else if err := validation.Validate(r.Sku, validation.Required); err != nil {
		return errs.NewBadRequestError("SKU is required")
	} else if err := validation.Validate(r.ProductCategoryID, validation.Required); err != nil {
		return errs.NewBadRequestError("Product category ID is required")
	} else if err := validation.Validate(r.Price, validation.Required); err != nil {
		return errs.NewBadRequestError("Price is required")
	} else if r.Price < 100 {
		return errs.NewValidationError("Minimum price is 100")
	} else if len(r.ProductCategoryID) < 1 {
		return errs.NewValidationError("Product category ID required")
	}
	return nil
}

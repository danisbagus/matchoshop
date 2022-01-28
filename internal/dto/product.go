package dto

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateProductRequest struct {
	MerchantID        int64   `json:"-"`
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
	MerchantID        int64    `json:"merchant_id"`
	Name              string   `json:"name"`
	Sku               string   `json:"sku"`
	Price             int64    `json:"price"`
	ProductCategories []string `json:"product_categories"`
}

type ProductListListResponse struct {
	Products []ProductListResponse `json:"data"`
}

type ProductDetailtResponse struct {
	ProductID         int64                     `json:"product_id"`
	MerchantID        int64                     `json:"merchant_id"`
	Name              string                    `json:"name"`
	Sku               string                    `json:"sku"`
	Price             int64                     `json:"price"`
	Description       string                    `json:"description"`
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
				MerchantID:        data[i].MerchantID,
				Name:              data[i].Name,
				Sku:               data[i].Sku,
				Price:             data[i].Price,
				ProductCategories: []string{data[i].ProductCategoryName},
			}
			mapProduct[data[i].ProductID] = product
		}
	}

	for _, valData := range mapProduct {
		products = append(products, valData)
	}

	productListResponse := &ProductListListResponse{Products: products}

	return GenerateResponseData(message, productListResponse)
}

func NewGetProductDetailResponse(message string, data *domain.ProductDetail) *ResponseData {
	productDetail := &ProductDetailtResponse{
		ProductID:   data.ProductID,
		MerchantID:  data.MerchantID,
		Name:        data.Name,
		Sku:         data.Sku,
		Price:       data.Price,
		Description: data.Description,
	}

	productCategories := make([]ProductCategoryResponse, 0)

	for _, valData := range data.ProductCategories {
		productCategory := ProductCategoryResponse{
			ProductCategoryID: valData.ProductCategoryID,
			MerchantID:        valData.MerchantID,
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

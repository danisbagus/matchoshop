package dto

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateProductCategoryRequest struct {
	MerchantID int64  `json:"merchant_id"`
	Name       string `json:"name"`
}

type CreateProductCategoryResponse struct {
	ProductCategoryID int64 `json:"product_category_id"`
}

type ProductCategoryResponse struct {
	ProductCategoryID int64  `json:"product_category_id"`
	MerchantID        int64  `json:"merchant_id"`
	Name              string `json:"name"`
}

type ProductCategoryListResponse struct {
	ProductCategories []ProductCategoryResponse `json:"data"`
}

func NewCreateProductCategoryResponse(data *domain.ProductCategory) *CreateProductCategoryResponse {
	result := &CreateProductCategoryResponse{
		ProductCategoryID: data.ProductCategoryID,
	}

	return result
}

func NewGetProductCategoryListResponse(data []domain.ProductCategory) *ProductCategoryListResponse {

	productCategories := make([]ProductCategoryResponse, len(data))

	for keyData, valData := range data {
		productCategories[keyData] = ProductCategoryResponse{
			ProductCategoryID: valData.ProductCategoryID,
			MerchantID:        valData.MerchantID,
			Name:              valData.Name,
		}
	}

	return &ProductCategoryListResponse{ProductCategories: productCategories}
}

func (r CreateProductCategoryRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.MerchantID, validation.Required); err != nil {
		return errs.NewBadRequestError("Merchant ID is required")

	}

	if err := validation.Validate(r.Name, validation.Required); err != nil {
		return errs.NewBadRequestError("Product category name is required")

	}

	return nil
}

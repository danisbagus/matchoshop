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

func NewCreateProductCategoryResponse(data *domain.ProductCategory) *CreateProductCategoryResponse {
	result := &CreateProductCategoryResponse{
		ProductCategoryID: data.ProductCategoryID,
	}

	return result
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

package dto

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateProductRequest struct {
	MerchantID        int64  `json:"-"`
	Name              string `json:"name"`
	Sku               string `json:"sku"`
	Description       string `json:"description"`
	ProductCategoryID int64  `json:"product_category_id"`
	Price             int64  `json:"price"`
}

type CreateProductResponse struct {
	ProductID int64 `json:"product_id"`
}

func NewCreateProductResponse(message string, data *domain.Product) *ResponseData {

	createProductResponse := &CreateProductResponse{
		ProductID: data.ProductID,
	}

	return GenerateResponseData(message, createProductResponse)
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
	}

	return nil
}

package domain

import (
	"github.com/danisbagus/go-common-packages/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	ProductCategoryModel struct {
		ProductCategoryID int64  `db:"product_category_id"`
		Name              string `db:"name"`
		CreatedAt         string `db:"created_at"`
		UpdatedAt         string `db:"updated_at"`
	}

	CreateProductCategoryRequest struct {
		Name string `json:"name"`
	}

	CreateProductCategoryResponse struct {
		ProductCategoryID int64 `json:"product_category_id"`
	}

	ProductCategoryResponse struct {
		ProductCategoryID int64  `json:"product_category_id"`
		Name              string `json:"name"`
	}

	UpdateroductCategoryRequest struct {
		Name string `json:"name"`
	}
)

func NewCreateProductCategoryResponse(message string, data *ProductCategoryModel) *ResponseData {

	createProductCategoryResponse := &CreateProductCategoryResponse{
		ProductCategoryID: data.ProductCategoryID,
	}

	return GenerateResponseData(message, createProductCategoryResponse)
}

func NewGetProductCategoryListResponse(message string, data []ProductCategoryModel) *ResponseData {

	productCategories := make([]ProductCategoryResponse, len(data))

	for keyData, valData := range data {
		productCategories[keyData] = ProductCategoryResponse{
			ProductCategoryID: valData.ProductCategoryID,
			Name:              valData.Name,
		}
	}
	return GenerateResponseData(message, productCategories)
}

func NewGetProductCategoryDetailResponse(message string, data *ProductCategoryModel) *ResponseData {
	productCategoryResponse := &ProductCategoryResponse{
		ProductCategoryID: data.ProductCategoryID,
		Name:              data.Name,
	}

	return GenerateResponseData(message, productCategoryResponse)
}

func (r CreateProductCategoryRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Name, validation.Required); err != nil {
		return errs.NewBadRequestError("Product category name is required")

	}

	return nil
}

func (r UpdateroductCategoryRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Name, validation.Required); err != nil {
		return errs.NewBadRequestError("Product category name is required")

	}

	return nil
}

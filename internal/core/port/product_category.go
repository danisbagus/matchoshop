package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/domain"
)

type ProductCategoryService interface {
	Create(data *domain.CreateProductCategoryRequest) (*domain.ResponseData, *errs.AppError)
	GetList() ([]domain.ProductCategoryModel, *errs.AppError)
	GetDetail(productCategoryID int64) (*domain.ResponseData, *errs.AppError)
	Update(productCategoryID int64, data *domain.CreateProductCategoryRequest) (*domain.ResponseData, *errs.AppError)
	Delete(productCategoryID int64) (*domain.ResponseData, *errs.AppError)
}

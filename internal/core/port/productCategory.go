package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type ProductCategoryService interface {
	Create(data *dto.CreateProductCategoryRequest) (*dto.ResponseData, *errs.AppError)
	GetList() ([]domain.ProductCategory, *errs.AppError)
	GetDetail(productCategoryID int64) (*dto.ResponseData, *errs.AppError)
	Update(productCategoryID int64, data *dto.CreateProductCategoryRequest) (*dto.ResponseData, *errs.AppError)
	Delete(productCategoryID int64) (*dto.ResponseData, *errs.AppError)
}

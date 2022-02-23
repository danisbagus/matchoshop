package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type ProductCategoryRepo interface {
	Insert(data *domain.ProductCategory) (*domain.ProductCategory, *errs.AppError)
	CheckByIDAndName(productCategoryID int64, name string) (bool, *errs.AppError)
	CheckByName(name string) (bool, *errs.AppError)
	CheckByID(productCategoryID int64) (bool, *errs.AppError)
	GetAll() ([]domain.ProductCategory, *errs.AppError)
	GetAllByProductID(productID int64) ([]domain.ProductCategory, *errs.AppError)
	GetOneByID(productCategoryID int64) (*domain.ProductCategory, *errs.AppError)
	Update(productCategoryID int64, data *domain.ProductCategory) *errs.AppError
	Delete(productCategoryID int64) *errs.AppError
}

type ProductCategoryService interface {
	Create(data *dto.CreateProductCategoryRequest) (*dto.ResponseData, *errs.AppError)
	GetList() (*dto.ResponseData, *errs.AppError)
	GetDetail(productCategoryID int64) (*dto.ResponseData, *errs.AppError)
	Update(productCategoryID int64, data *dto.CreateProductCategoryRequest) (*dto.ResponseData, *errs.AppError)
	Delete(productCategoryID int64) (*dto.ResponseData, *errs.AppError)
}

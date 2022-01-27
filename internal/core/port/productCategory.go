package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type IProductCategoryRepo interface {
	Insert(data *domain.ProductCategory) (*domain.ProductCategory, *errs.AppError)
	CheckByMerchantIDAndName(merchantID int64, name string) (bool, *errs.AppError)
	GetAllByMerchantID(merchantID int64) ([]domain.ProductCategory, *errs.AppError)
	GetOneByIDAndMerchantID(productCategoryID int64, merchantID int64) (*domain.ProductCategory, *errs.AppError)
}

type IProductCategoryService interface {
	Create(data *dto.CreateProductCategoryRequest) (*dto.CreateProductCategoryResponse, *errs.AppError)
	GetList(merchantID int64) (*dto.ProductCategoryListResponse, *errs.AppError)
	GetDetail(productCategoryID int64, merchantID int64) (*dto.ProductCategoryResponse, *errs.AppError)
}

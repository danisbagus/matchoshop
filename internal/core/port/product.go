package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type IProductRepo interface {
	Insert(data *domain.Product) (*domain.Product, *errs.AppError)
	CheckBySKUAndMerchantID(sku string, merchantID int64) (bool, *errs.AppError)
}

type IProductService interface {
	Create(data *dto.CreateProductRequest) (*dto.ResponseData, *errs.AppError)
}

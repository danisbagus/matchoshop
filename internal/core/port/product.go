package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/domain"
)

type ProductService interface {
	Create(form *domain.Product) *errs.AppError
	GetList(criteria *domain.ProductListCriteria) ([]domain.ProductDetail, *errs.AppError)
	GetListPaginate(criteria *domain.ProductListCriteria) ([]domain.ProductDetail, int64, *errs.AppError)
	GetDetail(productID int64) (*domain.ProductDetail, *errs.AppError)
	Update(productID int64, form *domain.Product) *errs.AppError
	Delete(productID int64) *errs.AppError
}

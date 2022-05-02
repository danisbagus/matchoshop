package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
)

type ProductRepo interface {
	Insert(data *domain.Product) (*domain.Product, *errs.AppError)
	CheckByID(productID int64) (bool, *errs.AppError)
	CheckBySKU(sku string) (bool, *errs.AppError)
	CheckByIDAndSKU(productID int64, sku string) (bool, *errs.AppError)
	GetAll() ([]domain.ProductList, *errs.AppError)
	GetOneByID(productID int64) (*domain.ProductDetail, *errs.AppError)
	Update(productID int64, data *domain.Product) *errs.AppError
	UpdateStock(productID, quantity int64) *errs.AppError
	Delete(productID int64) *errs.AppError
}

type ProductService interface {
	Create(form *domain.Product) *errs.AppError
	GetList() ([]domain.ProductDetail, *errs.AppError)
	GetDetail(productID int64) (*domain.ProductDetail, *errs.AppError)
	Update(productID int64, form *domain.Product) *errs.AppError
	Delete(productID int64) *errs.AppError
}

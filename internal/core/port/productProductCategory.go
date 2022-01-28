package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
)

type IProductProductCategoryRepo interface {
	BulkInsert(data []domain.ProductProductCategory) *errs.AppError
	DeleteAll(productID int64) *errs.AppError
}

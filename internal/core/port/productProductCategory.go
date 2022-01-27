package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
)

type IProductProductCategoryRepo interface {
	Insert(data *domain.ProductProductCategory) *errs.AppError
}

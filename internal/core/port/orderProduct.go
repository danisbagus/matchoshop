package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
)

type OrderProductRepo interface {
	GetAllByOrderID(orderUD int64) ([]domain.OrderProduct, *errs.AppError)
}

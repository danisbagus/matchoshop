package port

import (
	"github.com/danisbagus/go-common-packages/errs"
)

type (
	PaymentResultRepo interface {
		CheckByID(PaymentResultID string) (bool, *errs.AppError)
		CheckByOrderIDAndStatus(OrderID int64, status string) (bool, *errs.AppError)
	}
)

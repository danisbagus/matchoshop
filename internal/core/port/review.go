package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/domain"
)

type (
	ReviewService interface {
		Create(form *domain.Review) *errs.AppError
		GetDetail(userID, productID int64) (*domain.Review, *errs.AppError)
		Update(form *domain.Review) *errs.AppError
	}
)

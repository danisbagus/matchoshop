package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
)

type (
	ReviewRepo interface {
		Insert(form *domain.Review) *errs.AppError
		GetAllByProductID(productID int64) ([]domain.Review, *errs.AppError)
		GetOneByUserIDAndProductID(userID, productID int64) (*domain.Review, *errs.AppError)
		Update(form *domain.Review) *errs.AppError
	}

	ReviewService interface {
		Create(form *domain.Review) *errs.AppError
		GetDetail(userID, productID int64) (*domain.Review, *errs.AppError)
		Update(form *domain.Review) *errs.AppError
	}
)

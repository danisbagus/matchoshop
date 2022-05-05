package service

import (
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
)

type (
	ReviewService struct {
		repo port.ReviewRepo
	}
)

func NewReviewService(repo port.ReviewRepo) port.ReviewService {
	return &ReviewService{
		repo: repo,
	}
}

func (s ReviewService) GetDetail(userID, productID int64) (*domain.Review, *errs.AppError) {
	review, appErr := s.repo.GetOneByUserIDAndProductID(userID, productID)
	if appErr != nil {
		return nil, appErr
	}
	return review, nil
}

func (s ReviewService) Create(form *domain.Review) *errs.AppError {
	review, appErr := s.repo.GetOneByUserIDAndProductID(form.UserID, form.ProductID)
	if appErr != nil {
		return appErr
	}
	if review.ReviewID != 0 {
		return errs.NewBadRequestError("review already created")
	}

	form.CreatedAt = time.Now().Format(dbTSLayout)
	form.UpdatedAt = time.Now().Format(dbTSLayout)
	appErr = s.repo.Insert(form)
	if appErr != nil {
		return appErr
	}

	return nil
}

func (s ReviewService) Update(form *domain.Review) *errs.AppError {
	review, appErr := s.repo.GetOneByUserIDAndProductID(form.UserID, form.ProductID)
	if appErr != nil {
		return appErr
	}
	if review.ReviewID == 0 {
		return errs.NewBadRequestError("review not found")
	}

	form.ReviewID = review.ReviewID
	form.UpdatedAt = time.Now().Format(dbTSLayout)
	appErr = s.repo.Update(form)
	if appErr != nil {
		return appErr
	}

	return nil
}

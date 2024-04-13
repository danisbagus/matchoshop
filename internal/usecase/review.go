package usecase

import (
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/danisbagus/matchoshop/internal/repository"
)

type IReviewUsecase interface {
	Create(form *domain.Review) *errs.AppError
	GetDetail(userID, productID int64) (*domain.Review, *errs.AppError)
	Update(form *domain.Review) *errs.AppError
}

type ReviewUsecase struct {
	reviewRepo repository.IReviewRepository
}

func NewReviewUsecase(repository repository.RepositoryCollection) IReviewUsecase {
	return &ReviewUsecase{
		reviewRepo: repository.ReviewRepository,
	}
}

func (s ReviewUsecase) GetDetail(userID, productID int64) (*domain.Review, *errs.AppError) {
	review, appErr := s.reviewRepo.GetOneByUserIDAndProductID(userID, productID)
	if appErr != nil {
		return nil, appErr
	}
	return review, nil
}

func (s ReviewUsecase) Create(form *domain.Review) *errs.AppError {
	review, appErr := s.reviewRepo.GetOneByUserIDAndProductID(form.UserID, form.ProductID)
	if appErr != nil {
		return appErr
	}
	if review.ReviewID != 0 {
		return errs.NewBadRequestError("review already created")
	}

	form.CreatedAt = time.Now().Format(dbTSLayout)
	form.UpdatedAt = time.Now().Format(dbTSLayout)
	appErr = s.reviewRepo.Insert(form)
	if appErr != nil {
		return appErr
	}

	return nil
}

func (s ReviewUsecase) Update(form *domain.Review) *errs.AppError {
	review, appErr := s.reviewRepo.GetOneByUserIDAndProductID(form.UserID, form.ProductID)
	if appErr != nil {
		return appErr
	}
	if review.ReviewID == 0 {
		return errs.NewBadRequestError("review not found")
	}

	form.ReviewID = review.ReviewID
	form.UpdatedAt = time.Now().Format(dbTSLayout)
	appErr = s.reviewRepo.Update(form)
	if appErr != nil {
		return appErr
	}

	return nil
}

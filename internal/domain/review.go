package domain

import (
	"github.com/danisbagus/go-common-packages/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	ReviewModel struct {
		ReviewID  int64  `db:"review_id"`
		UserID    int64  `db:"user_id"`
		ProductID int64  `db:"product_id"`
		Rating    int    `db:"rating"`
		Comment   string `db:"comment"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}

	Review struct {
		ReviewModel
		UserName string
	}

	ReviewRequest struct {
		ReviewID  int64  `json:"-"`
		UserID    int64  `json:"-"`
		ProductID int64  `json:"product_id"`
		Rating    int    `json:"rating"`
		Comment   string `json:"comment"`
	}

	ReviewResponse struct {
		ReviewID  int64  `json:"review_id"`
		UserID    int64  `json:"user_id"`
		UserName  string `json:"user_name"`
		ProductID int64  `json:"product_id"`
		Rating    int    `json:"rating"`
		Comment   string `json:"comment"`
		CreatedAt string `json:"created_at"`
	}
)

func NewReviewResponse(message string, data *Review) *ResponseData {
	resData := new(ReviewResponse)
	resData.ReviewID = data.ReviewID
	resData.UserID = data.UserID
	resData.ProductID = data.ProductID
	resData.Rating = data.Rating
	resData.Comment = data.Comment
	resData.CreatedAt = data.CreatedAt

	return GenerateResponseData(message, resData)
}

func (r ReviewRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.ProductID, validation.Required); err != nil {
		return errs.NewBadRequestError("product id is required")
	} else if err := validation.Validate(r.Rating, validation.Required); err != nil {
		return errs.NewBadRequestError("rating is required")
	} else if err := validation.Validate(r.Rating, validation.Min(1)); err != nil {
		return errs.NewBadRequestError("minimal rating value is 1")
	} else if err := validation.Validate(r.Rating, validation.Max(5)); err != nil {
		return errs.NewBadRequestError("maximal rating value is 5")
	} else if err := validation.Validate(r.Comment, validation.Required); err != nil {
		return errs.NewBadRequestError("comment is required")
	}
	return nil
}

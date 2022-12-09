package repo

import (
	"database/sql"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
)

type ReviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) port.ReviewRepo {
	return &ReviewRepo{
		db: db,
	}
}

func (r ReviewRepo) GetAllByProductID(productID int64) ([]domain.Review, *errs.AppError) {
	sqlGet := `
	SELECT 
		r.review_id, 
		r.user_id, 
		r.product_id, 
		r.rating, 
		r.comment,
		r.created_at,
		u.name
	FROM 
		reviews r 
	INNER JOIN users u ON u.user_id = r.user_id
	WHERE 
		r.product_id=$1`

	rows, err := r.db.Query(sqlGet, productID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get all reviews from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	reviews := make([]domain.Review, 0)
	for rows.Next() {
		var review domain.Review
		err := rows.Scan(&review.ReviewID, &review.UserID, &review.ProductID, &review.Rating, &review.Comment, &review.CreatedAt, &review.UserName)
		if err != nil && err != sql.ErrNoRows {
			logger.Error("Error while reviews all order from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r ReviewRepo) GetOneByUserIDAndProductID(UserID, ProductID int64) (*domain.Review, *errs.AppError) {
	sqlGet := `
	SELECT 
		r.review_id, 
		r.user_id, 
		r.product_id, 
		r.rating, 
		r.comment,
		r.created_at
	FROM 
		reviews r 
	WHERE 
  		r.user_id=$1
		AND r.product_id=$2`

	var review domain.Review
	err := r.db.QueryRow(sqlGet, UserID, ProductID).Scan(&review.ReviewID, &review.UserID, &review.ProductID, &review.Rating, &review.Comment, &review.CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get order from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")

	}

	return &review, nil
}

func (r ReviewRepo) Insert(form *domain.Review) *errs.AppError {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting insert review: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInsert := `INSERT INTO reviews(user_id, product_id, rating, comment, created_at, updated_at) 
					  VALUES($1, $2,$3, $4, $5, $6)`

	err = tx.QueryRow(sqlInsert, form.UserID, form.ProductID, form.Rating, form.Comment, form.CreatedAt, form.UpdatedAt).Err()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert review: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}

func (r ReviewRepo) Update(form *domain.Review) *errs.AppError {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting update review: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlUpdate := `
	UPDATE reviews 
	SET rating = $2,
		comment = $3,
		updated_at = $4
	WHERE review_id = $1`

	_, err = tx.Exec(sqlUpdate, form.ReviewID, form.Rating, form.Comment, time.Now())
	if err != nil {
		tx.Rollback()
		logger.Error("Error while update review: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}

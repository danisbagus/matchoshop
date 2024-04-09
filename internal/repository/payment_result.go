package repository

import (
	"database/sql"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/jmoiron/sqlx"
)

type IPaymentResultRepository interface {
	CheckByID(PaymentResultID string) (bool, *errs.AppError)
	CheckByOrderIDAndStatus(OrderID int64, status string) (bool, *errs.AppError)
}

type PaymentResultRepository struct {
	db *sqlx.DB
}

func NewPaymentResultRepository(db *sqlx.DB) *PaymentResultRepository {
	return &PaymentResultRepository{
		db: db,
	}
}

func (r PaymentResultRepository) CheckByID(ID string) (bool, *errs.AppError) {
	sqlGet := `
	SELECT 
	    payment_result_id
	FROM 
		payment_results
	WHERE 
	payment_result_id=$1`

	var paymentResultID string
	err := r.db.QueryRow(sqlGet, ID).Scan(&paymentResultID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while check payment result from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")

	}
	return paymentResultID != "", nil
}

func (r PaymentResultRepository) CheckByOrderIDAndStatus(OrderID int64, status string) (bool, *errs.AppError) {
	sqlGet := `
	SELECT 
	    payment_result_id
	FROM 
		payment_results
	WHERE 
  		order_id=$1 AND status=$2`

	var paymentResultID string
	err := r.db.QueryRow(sqlGet, OrderID, status).Scan(&paymentResultID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while check payment result from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")

	}
	return paymentResultID != "", nil
}

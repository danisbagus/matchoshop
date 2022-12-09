package repo

import (
	"database/sql"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
)

type PaymentResult struct {
	db *sql.DB
}

func NewPaymentResult(db *sql.DB) port.PaymentResultRepo {
	return &PaymentResult{
		db: db,
	}
}

func (r PaymentResult) CheckByID(ID string) (bool, *errs.AppError) {
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

func (r PaymentResult) CheckByOrderIDAndStatus(OrderID int64, status string) (bool, *errs.AppError) {
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

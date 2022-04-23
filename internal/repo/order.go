package repo

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) port.OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r OrderRepo) GetAllByUserID(userID int64) ([]domain.OrderDetail, *errs.AppError) {
	sqlGet := `
	SELECT 
		o.order_id, 
		o.user_id, 
		o.payment_method_id, 
		o.product_price, 
		o.tax_price, 
		o.shipping_price, 
		o.total_price, 
		o.is_paid, 
		o.paid_at, 
		o.is_delivered,
		o.delivered_at,
		o.created_at
	FROM 
		orders o 
	WHERE 
  		o.user_id=$1`

	rows, err := r.db.Query(sqlGet, userID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get all order from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	orders := make([]domain.OrderDetail, 0)
	for rows.Next() {
		var order domain.OrderDetail
		err := rows.Scan(&order.Order.OrderID, &order.UserID, &order.PaymentMethodID, &order.ProductPrice, &order.TaxPrice, &order.ShippingPrice,
			&order.TotalPrice, &order.IsPaid, &order.PaidAt, &order.IsDelivered, &order.DeliveredAt, &order.CreatedAt)
		if err != nil && err != sql.ErrNoRows {
			logger.Error("Error while get all order from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r OrderRepo) GetOneByID(OrderID int64) (*domain.OrderDetail, *errs.AppError) {
	sqlGet := `
	SELECT 
		o.order_id, 
		o.user_id, 
		o.payment_method_id, 
		o.product_price, 
		o.tax_price, 
		o.shipping_price, 
		o.total_price, 
		o.is_paid, 
		o.paid_at, 
		sa.address, 
		sa.city, 
		sa.postal_code, 
		sa.country,
		pm.name,
		u.name AS user_name,
		u.email AS user_email
	FROM 
		orders o 
		LEFT JOIN shipment_address sa ON sa.order_id = o.order_id
		INNER JOIN payment_methods pm ON pm.payment_method_id = o.payment_method_id
		INNER JOIN users u ON u.user_id = o.user_id
	WHERE 
  		o.order_id=$1`

	var order domain.OrderDetail
	err := r.db.QueryRow(sqlGet, OrderID).Scan(&order.Order.OrderID, &order.UserID, &order.PaymentMethodID, &order.ProductPrice, &order.TaxPrice, &order.ShippingPrice,
		&order.TotalPrice, &order.IsPaid, &order.PaidAt, &order.Address, &order.City, &order.PostalCode, &order.Country, &order.PaymentMethodName, &order.UserName, &order.UserEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Order not found!")
		} else {
			logger.Error("Error while get order from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &order, nil
}

func (r OrderRepo) Insert(form *domain.OrderDetail) (int64, *errs.AppError) {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting insert order: " + err.Error())
		return 0, errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInsert := `INSERT INTO orders(user_id, payment_method_id, product_price, tax_price, shipping_price, total_price, created_at, updated_at) 
					  VALUES($1, $2,$3, $4, $5, $6, $7, $8)
					  RETURNING order_id`

	var orderID int64
	err = tx.QueryRow(sqlInsert, form.UserID, form.PaymentMethodID, form.ProductPrice, form.TaxPrice, form.ShippingPrice, form.TotalPrice, form.CreatedAt, form.UpdatedAt).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert order: " + err.Error())
		return 0, errs.NewUnexpectedError("Unexpected database error")
	}

	err = r.bulkInsertOrderProduct(tx, orderID, form.OrderProducts)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert order product: " + err.Error())
		return 0, errs.NewUnexpectedError("Unexpected database error")
	}

	err = r.insertShipmentAddress(tx, orderID, &form.ShipmentAddress)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert order product: " + err.Error())
		return 0, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return 0, errs.NewUnexpectedError("Unexpected database error")
	}

	return orderID, nil
}

func (r OrderRepo) UpdatePaid(form *domain.PaymentResult) *errs.AppError {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting update order paid: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlUpdate := `
	UPDATE orders 
	SET is_paid = $2,
		paid_at = $3,
		updated_at = $3
	WHERE order_id = $1`

	_, err = tx.Exec(sqlUpdate, form.OrderID, 1, time.Now())
	if err != nil {
		tx.Rollback()
		logger.Error("Error while update product: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	err = r.insertPaymentResult(tx, form)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert order product: " + err.Error())
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

func (r OrderRepo) bulkInsertOrderProduct(tx *sql.Tx, orderID int64, form []domain.OrderProduct) error {
	valueStrings := make([]string, 0, len(form))
	valueArgs := make([]interface{}, 0, len(form)*2)

	i := 0
	for _, post := range form {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
		valueArgs = append(valueArgs, orderID)
		valueArgs = append(valueArgs, post.ProductID)
		valueArgs = append(valueArgs, post.Quantity)
		i++
	}

	sqlInsert := fmt.Sprintf("INSERT INTO order_products (order_id, product_id, quantity) VALUES %s",
		strings.Join(valueStrings, ","))

	_, err := tx.Exec(sqlInsert, valueArgs...)
	if err != nil {
		return err
	}
	return nil
}

func (r OrderRepo) insertShipmentAddress(tx *sql.Tx, orderID int64, form *domain.ShipmentAddress) error {

	sqlInsert := `INSERT INTO shipment_address(order_id, address, city, postal_code, country) 
					  VALUES($1, $2,$3, $4, $5)`

	_, err := tx.Exec(sqlInsert, orderID, form.Address, form.City, form.PostalCode, form.Country)
	if err != nil {
		return err
	}
	return nil
}

func (r OrderRepo) insertPaymentResult(tx *sql.Tx, form *domain.PaymentResult) error {

	sqlInsert := `INSERT INTO payment_results(payment_result_id, order_id, status, update_time, email) 
					  VALUES($1, $2, $3, $4, $5)`

	_, err := tx.Exec(sqlInsert, form.PaymentResultID, form.OrderID, form.Status, form.UpdateTime, form.Email)
	if err != nil {
		return err
	}
	return nil
}

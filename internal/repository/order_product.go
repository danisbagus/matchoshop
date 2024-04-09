package repository

import (
	"database/sql"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/jmoiron/sqlx"
)

type IOrderProductRepository interface {
	GetAllByOrderID(orderUD int64) ([]domain.OrderProduct, *errs.AppError)
}

type OrderProductRepository struct {
	db *sqlx.DB
}

func NewOrderProductRepository(db *sqlx.DB) *OrderProductRepository {
	return &OrderProductRepository{
		db: db,
	}
}

func (r *OrderProductRepository) GetAllByOrderID(orderID int64) ([]domain.OrderProduct, *errs.AppError) {

	sqlGet := `
	SELECT 
		op.order_id, 
		op.product_id, 
		p.name, 
		p.price, 
		p.image,
		op.quantity 
	FROM 
		order_products op 
		INNER JOIN products p ON p.product_id = op.product_id 
	WHERE 
		op.order_id=$1
  `
	rows, err := r.db.Query(sqlGet, orderID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get all order product from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	orderProducts := make([]domain.OrderProduct, 0)
	for rows.Next() {
		var orderProduct domain.OrderProduct
		if err := rows.Scan(&orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Name, &orderProduct.Price, &orderProduct.Image, &orderProduct.Quantity); err != nil {
			logger.Error("Error while scanning porder productfrom database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		orderProducts = append(orderProducts, orderProduct)
	}

	return orderProducts, nil
}

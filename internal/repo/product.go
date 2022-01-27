package repo

import (
	"database/sql"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) port.IProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r ProductRepo) Insert(data *domain.Product) (*domain.Product, *errs.AppError) {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting insert product: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInsert := `INSERT INTO products(merchant_id, name, sku, description, price, created_at, updated_at) 
					  VALUES($1, $2, $3, $4, $5, $6, $7)
					  RETURNING product_id`

	var productID int64
	err = tx.QueryRow(sqlInsert, data.MerchantID, data.Name, data.Sku, data.Description, data.Price, data.CreatedAt, data.UpdatedAt).Scan(&productID)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert product: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	data.ProductID = productID

	return data, nil
}

func (r ProductRepo) CheckBySKUAndMerchantID(sku string, merchantID int64) (bool, *errs.AppError) {

	sqlCountProduct := `SELECT COUNT(product_Id) 
	FROM products 
	WHERE merchant_id = $1
	AND sku = $2`

	var totalData int64
	err := r.db.QueryRow(sqlCountProduct, merchantID, sku).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

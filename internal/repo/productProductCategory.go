package repo

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/jmoiron/sqlx"
)

type ProductProductCategoryRepo struct {
	db *sqlx.DB
}

func NewProductProductCategoryRepo(db *sqlx.DB) port.IProductProductCategoryRepo {
	return &ProductProductCategoryRepo{
		db: db,
	}
}

func (r ProductProductCategoryRepo) Insert(data *domain.ProductProductCategory) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting insert product product category: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInsert := `INSERT INTO product_product_categories(product_id, product_category_id) VALUES($1, $2)`

	_, err = tx.Exec(sqlInsert, data.ProductID, data.ProductCategoryID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert product product category: " + err.Error())
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
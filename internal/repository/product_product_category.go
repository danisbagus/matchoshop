package repository

import (
	"fmt"
	"strings"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/jmoiron/sqlx"
)

type IProductProductCategoryRepository interface {
	BulkInsert(data []domain.ProductProductCategory) *errs.AppError
	DeleteAll(productID int64) *errs.AppError
}

type ProductProductCategoryRepository struct {
	db *sqlx.DB
}

func NewProductProductCategoryRepository(db *sqlx.DB) *ProductProductCategoryRepository {
	return &ProductProductCategoryRepository{
		db: db,
	}
}

func (r ProductProductCategoryRepository) BulkInsert(data []domain.ProductProductCategory) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting insert product product category: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	valueStrings := make([]string, 0, len(data))
	valueArgs := make([]interface{}, 0, len(data)*2)

	i := 0
	for _, post := range data {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, post.ProductID)
		valueArgs = append(valueArgs, post.ProductCategoryID)
		i++
	}

	sqlInsert := fmt.Sprintf("INSERT INTO product_product_categories (product_id, product_category_id) VALUES %s",
		strings.Join(valueStrings, ","))

	_, err = tx.Exec(sqlInsert, valueArgs...)
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

func (r ProductProductCategoryRepository) DeleteAll(productID int64) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting delete product product category: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlDeleteAll := `
	DELETE FROM product_product_categories 
	WHERE product_id = $1`

	_, err = tx.Exec(sqlDeleteAll, productID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while delete product product category: " + err.Error())
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

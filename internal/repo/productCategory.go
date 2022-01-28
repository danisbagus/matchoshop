package repo

import (
	"database/sql"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/jmoiron/sqlx"
)

type ProductCategoryRepo struct {
	db *sqlx.DB
}

func NewProductCategoryRepo(db *sqlx.DB) port.IProductCategoryRepo {
	return &ProductCategoryRepo{
		db: db,
	}
}

func (r ProductCategoryRepo) Insert(data *domain.ProductCategory) (*domain.ProductCategory, *errs.AppError) {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting insert product category: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInsert := `INSERT INTO product_categories(merchant_id, name, created_at, updated_at) 
					  VALUES($1, $2, $3, $4)
					  RETURNING product_category_id`

	var productCategoryID int64
	err = tx.QueryRow(sqlInsert, data.MerchantID, data.Name, data.CreatedAt, data.UpdatedAt).Scan(&productCategoryID)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while insert product category: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	data.ProductCategoryID = productCategoryID

	return data, nil
}

func (r ProductCategoryRepo) CheckByIDAndMerchantIDAndName(productCategoryID int64, merchantID int64, name string) (bool, *errs.AppError) {

	sqlCountProductCategory := `SELECT COUNT(product_category_Id) 
	FROM product_categories 
	WHERE product_category_Id != $1
	AND merchant_id = $2
	AND name = $3`

	var totalData int64
	err := r.db.QueryRow(sqlCountProductCategory, productCategoryID, merchantID, name).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product category from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductCategoryRepo) CheckByMerchantIDAndName(merchantID int64, name string) (bool, *errs.AppError) {

	sqlCountProductCategory := `SELECT COUNT(product_category_Id) 
	FROM product_categories 
	WHERE merchant_id = $1
	AND name = $2`

	var totalData int64
	err := r.db.QueryRow(sqlCountProductCategory, merchantID, name).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product category from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductCategoryRepo) CheckByIDAndMerchantID(productCategoryID int64, merchantID int64) (bool, *errs.AppError) {

	sqlCountProductCategory := `SELECT COUNT(product_category_Id) 
	FROM product_categories 
	WHERE product_category_id = $1
	AND merchant_id = $2`

	var totalData int64
	err := r.db.QueryRow(sqlCountProductCategory, productCategoryID, merchantID).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product category from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductCategoryRepo) GetAllByMerchantID(merchantID int64) ([]domain.ProductCategory, *errs.AppError) {

	sqlGetProductCategory := `
	SELECT 
		product_category_id, 
		merchant_id, 
		name
	FROM product_categories
	WHERE merchant_id = $1
	ORDER BY name ASC`

	rows, err := r.db.Query(sqlGetProductCategory, merchantID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get product category from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	productCategories := make([]domain.ProductCategory, 0)

	for rows.Next() {
		var productCategory domain.ProductCategory
		if err := rows.Scan(&productCategory.ProductCategoryID, &productCategory.MerchantID, &productCategory.Name); err != nil {
			logger.Error("Error while scanning product category from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		productCategories = append(productCategories, productCategory)
	}

	return productCategories, nil
}

func (r ProductCategoryRepo) GetOneByIDAndMerchantID(productCategoryID int64, merchantID int64) (*domain.ProductCategory, *errs.AppError) {

	var productCategory domain.ProductCategory

	sqlGetProductCategory := `
	SELECT 
		product_category_id, 
		merchant_id, 
		name
	FROM product_categories
	WHERE product_category_id = $1 
	AND merchant_id = $2
	LIMIT 1`

	err := r.db.QueryRow(sqlGetProductCategory, productCategoryID, merchantID).Scan(&productCategory.ProductCategoryID, &productCategory.MerchantID, &productCategory.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Product category not found!")
		} else {
			logger.Error("Error while get product category from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &productCategory, nil

}

func (r ProductCategoryRepo) Update(productCategoryID int64, data *domain.ProductCategory) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting update product category: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlUpdate := `
	UPDATE product_categories 
	SET merchant_id = $2, name = $3, created_at = $4, updated_at = $5
	WHERE product_category_id = $1`

	_, err = tx.Exec(sqlUpdate, productCategoryID, data.MerchantID, data.Name, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while update product category: " + err.Error())
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

func (r ProductCategoryRepo) Delete(productCategoryID int64) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting delete product category: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlDelete := `
	DELETE FROM product_categories 
	WHERE product_category_id = $1`

	_, err = tx.Exec(sqlDelete, productCategoryID)

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}

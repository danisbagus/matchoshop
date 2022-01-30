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

	sqlInsert := `INSERT INTO product_categories(name, created_at, updated_at) 
					  VALUES($1, $2, $3)
					  RETURNING product_category_id`

	var productCategoryID int64
	err = tx.QueryRow(sqlInsert, data.Name, data.CreatedAt, data.UpdatedAt).Scan(&productCategoryID)

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

func (r ProductCategoryRepo) CheckByIDAndName(productCategoryID int64, name string) (bool, *errs.AppError) {

	sqlCountProductCategory := `SELECT COUNT(product_category_Id) 
	FROM product_categories 
	WHERE product_category_Id != $1
	AND name = $2`

	var totalData int64
	err := r.db.QueryRow(sqlCountProductCategory, productCategoryID, name).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product category from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductCategoryRepo) CheckByName(name string) (bool, *errs.AppError) {

	sqlCountProductCategory := `SELECT COUNT(product_category_Id) 
	FROM product_categories 
	WHERE name = $1`

	var totalData int64
	err := r.db.QueryRow(sqlCountProductCategory, name).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product category from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductCategoryRepo) CheckByID(productCategoryID int64) (bool, *errs.AppError) {

	sqlCountProductCategory := `SELECT COUNT(product_category_Id) 
	FROM product_categories 
	WHERE product_category_id = $1`

	var totalData int64
	err := r.db.QueryRow(sqlCountProductCategory, productCategoryID).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product category from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductCategoryRepo) GetAll() ([]domain.ProductCategory, *errs.AppError) {

	sqlGetProductCategory := `
	SELECT 
		product_category_id, 
		name
	FROM product_categories
	ORDER BY name ASC`

	rows, err := r.db.Query(sqlGetProductCategory)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get product category from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	productCategories := make([]domain.ProductCategory, 0)

	for rows.Next() {
		var productCategory domain.ProductCategory
		if err := rows.Scan(&productCategory.ProductCategoryID, &productCategory.Name); err != nil {
			logger.Error("Error while scanning product category from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		productCategories = append(productCategories, productCategory)
	}

	return productCategories, nil
}

func (r ProductCategoryRepo) GetAllByProductID(productID int64) ([]domain.ProductCategory, *errs.AppError) {

	sqlGetProductCategory := `
	SELECT 
		pc.product_category_id, 
		pc.name
	FROM product_categories pc
	JOIN product_product_categories ppc ON ppc.product_category_id = pc.product_category_id
	WHERE ppc.product_id = $1`

	rows, err := r.db.Query(sqlGetProductCategory, productID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get product category from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	productCategories := make([]domain.ProductCategory, 0)

	for rows.Next() {
		var productCategory domain.ProductCategory
		if err := rows.Scan(&productCategory.ProductCategoryID, &productCategory.Name); err != nil {
			logger.Error("Error while scanning product category from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		productCategories = append(productCategories, productCategory)
	}

	return productCategories, nil
}

func (r ProductCategoryRepo) GetOneByID(productCategoryID int64) (*domain.ProductCategory, *errs.AppError) {

	var productCategory domain.ProductCategory

	sqlGetProductCategory := `
	SELECT 
		product_category_id, 
		name
	FROM product_categories
	WHERE product_category_id = $1 
	LIMIT 1`

	err := r.db.QueryRow(sqlGetProductCategory, productCategoryID).Scan(&productCategory.ProductCategoryID, &productCategory.Name)
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
	SET name = $2, created_at = $3, updated_at = $4
	WHERE product_category_id = $1`

	_, err = tx.Exec(sqlUpdate, productCategoryID, data.Name, data.CreatedAt, data.UpdatedAt)
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
	if err != nil {
		tx.Rollback()
		logger.Error("Error while delete product category: " + err.Error())
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

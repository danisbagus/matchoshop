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

	sqlInsert := `INSERT INTO products(name, sku, description, price, created_at, updated_at) 
					  VALUES($1, $2, $3, $4, $5, $6)
					  RETURNING product_id`

	var productID int64
	err = tx.QueryRow(sqlInsert, data.Name, data.Sku, data.Description, data.Price, data.CreatedAt, data.UpdatedAt).Scan(&productID)

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

func (r ProductRepo) CheckByID(productID int64) (bool, *errs.AppError) {

	sqlCountProduct := `SELECT COUNT(product_Id) 
	FROM products 
	WHERE product_id = $1`

	var totalData int64
	err := r.db.QueryRow(sqlCountProduct, productID).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductRepo) CheckBySKU(sku string) (bool, *errs.AppError) {

	sqlCountProduct := `SELECT COUNT(product_Id) 
	FROM products 
	WHERE sku = $1`

	var totalData int64
	err := r.db.QueryRow(sqlCountProduct, sku).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductRepo) CheckByIDAndSKU(productID int64, sku string) (bool, *errs.AppError) {

	sqlCountProduct := `SELECT COUNT(product_Id) 
	FROM products 
	WHERE product_id != $1
	AND sku = $2`

	var totalData int64
	err := r.db.QueryRow(sqlCountProduct, productID, sku).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductRepo) GetAll() ([]domain.ProductList, *errs.AppError) {

	sqlGetProduct := `
	SELECT 
		p.product_id, 
		p.name, 
		p.sku, 
		p.price, 
		pc.name as product_category_name
	FROM products p
	JOIN product_product_categories ppc ON ppc.product_id = p.product_id
	JOIN product_categories pc ON pc.product_category_id = ppc.product_category_id
	ORDER BY p.name ASC`

	rows, err := r.db.Query(sqlGetProduct)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get all product from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	products := make([]domain.ProductList, 0)

	for rows.Next() {
		var product domain.ProductList
		if err := rows.Scan(&product.ProductID, &product.Name, &product.Sku, &product.Price, &product.ProductCategoryName); err != nil {
			logger.Error("Error while scanning product category from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		products = append(products, product)
	}

	return products, nil
}

func (r ProductRepo) GetOneByID(productID int64) (*domain.ProductDetail, *errs.AppError) {

	var product domain.ProductDetail

	sqlGetProduct := `
	SELECT 
		p.product_id, 
		p.name, 
		p.sku, 
		p.price, 
		p.description
	FROM products p
	WHERE p.product_id = $1
	LIMIT 1`

	err := r.db.QueryRow(sqlGetProduct, productID).Scan(&product.ProductID, &product.Name, &product.Sku, &product.Price, &product.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Product not found!")
		} else {
			logger.Error("Error while get product from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &product, nil
}

func (r ProductRepo) Update(productID int64, data *domain.Product) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting update product: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlUpdate := `
	UPDATE products 
	SET name = $2, 
		sku = $3,
		price = $4,
		description = $5,
		created_at = $6, 
		updated_at = $7
	WHERE product_id = $1`

	_, err = tx.Exec(sqlUpdate, productID, data.Name, data.Sku, data.Price, data.Description, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while update product: " + err.Error())
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

func (r ProductRepo) Delete(productID int64) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting delete product: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlDelete := `
	DELETE FROM products 
	WHERE product_id = $1`

	_, err = tx.Exec(sqlDelete, productID)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while delete product: " + err.Error())
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

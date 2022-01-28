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

func (r ProductRepo) CheckByIDAndMerchantID(productID int64, merchantID int64) (bool, *errs.AppError) {

	sqlCountProduct := `SELECT COUNT(product_Id) 
	FROM products 
	WHERE product_id = $1
	AND merchant_id = $2`

	var totalData int64
	err := r.db.QueryRow(sqlCountProduct, productID, merchantID).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
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

func (r ProductRepo) CheckByIDAndSKUAndMerchantID(productID int64, sku string, merchantID int64) (bool, *errs.AppError) {

	sqlCountProduct := `SELECT COUNT(product_Id) 
	FROM products 
	WHERE product_id != $1
	AND merchant_id = $2
	AND sku = $3`

	var totalData int64
	err := r.db.QueryRow(sqlCountProduct, productID, merchantID, sku).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count product from database: " + err.Error())
		return false, errs.NewUnexpectedError("Unexpected database error")
	}

	return totalData > 0, nil
}

func (r ProductRepo) GetAllByMerchantID(merchantID int64) ([]domain.ProductList, *errs.AppError) {

	sqlGetProduct := `
	SELECT 
		p.product_id, 
		p.merchant_id, 
		p.name, 
		p.sku, 
		p.price, 
		pc.name as product_category_name
	FROM products p
	JOIN product_product_categories ppc ON ppc.product_id = p.product_id
	JOIN product_categories pc ON pc.product_category_id = ppc.product_category_id
	WHERE p.merchant_id = $1
	ORDER BY p.name ASC`

	rows, err := r.db.Query(sqlGetProduct, merchantID)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get all product from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	products := make([]domain.ProductList, 0)

	for rows.Next() {
		var product domain.ProductList
		if err := rows.Scan(&product.ProductID, &product.MerchantID, &product.Name, &product.Sku, &product.Price, &product.ProductCategoryName); err != nil {
			logger.Error("Error while scanning product category from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		products = append(products, product)
	}

	return products, nil
}

func (r ProductRepo) GetOneByIDAndMerchantID(productID int64, merchantID int64) (*domain.ProductDetail, *errs.AppError) {

	var product domain.ProductDetail

	sqlGetProduct := `
	SELECT 
		p.product_id, 
		p.merchant_id, 
		p.name, 
		p.sku, 
		p.price, 
		p.description
	FROM products p
	WHERE p.product_id = $1
	AND p.merchant_id = $2
	LIMIT 1`

	err := r.db.QueryRow(sqlGetProduct, productID, merchantID).Scan(&product.ProductID, &product.MerchantID, &product.Name, &product.Sku, &product.Price, &product.Description)
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
	SET merchant_id = $2, 
		name = $3, 
		sku = $4,
		price = $5,
		description = $6,
		created_at = $7, 
		updated_at = $8
	WHERE product_id = $1`

	_, err = tx.Exec(sqlUpdate, productID, data.MerchantID, data.Name, data.Sku, data.Price, data.Description, data.CreatedAt, data.UpdatedAt)
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

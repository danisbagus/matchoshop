package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/jmoiron/sqlx"
)

type IProductRepository interface {
	Insert(data *domain.Product) (*domain.Product, *errs.AppError)
	CheckByID(productID int64) (bool, *errs.AppError)
	CheckBySKU(sku string) (bool, *errs.AppError)
	CheckByIDAndSKU(productID int64, sku string) (bool, *errs.AppError)
	GetAll(criteria *domain.ProductListCriteria) ([]domain.ProductList, *errs.AppError)
	GetAllPaginate(criteria *domain.ProductListCriteria) ([]domain.ProductList, int64, *errs.AppError)
	GetOneByID(productID int64) (*domain.ProductDetail, *errs.AppError)
	Update(productID int64, data *domain.Product) *errs.AppError
	UpdateStock(productID, quantity int64) *errs.AppError
	Delete(productID int64) *errs.AppError
}

type ProductReposotory struct {
	db *sqlx.DB
}

func NewProductReposotory(db *sqlx.DB) *ProductReposotory {
	return &ProductReposotory{
		db: db,
	}
}

func (r ProductReposotory) Insert(data *domain.Product) (*domain.Product, *errs.AppError) {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting insert product: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	sqlInsert := `INSERT INTO products(name, sku, brand, image, description, price, stock, created_at, updated_at) 
					  VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
					  RETURNING product_id`

	var productID int64
	err = tx.QueryRow(sqlInsert, data.Name, data.Sku, data.Brand, data.Image, data.Description, data.Price, data.Stock, data.CreatedAt, data.UpdatedAt).Scan(&productID)

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

func (r ProductReposotory) CheckByID(productID int64) (bool, *errs.AppError) {

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

func (r ProductReposotory) CheckBySKU(sku string) (bool, *errs.AppError) {

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

func (r ProductReposotory) CheckByIDAndSKU(productID int64, sku string) (bool, *errs.AppError) {

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

func (r ProductReposotory) GetAll(criteria *domain.ProductListCriteria) ([]domain.ProductList, *errs.AppError) {
	sort := getSortProduct(criteria.Sort, criteria.Order)

	sqlGetProduct := fmt.Sprintf(`
	SELECT 
		p.product_id, 
		p.name, 
		p.sku, 
		p.brand, 
		p.image, 
		p.price, 
		pc.product_category_id,
		pc.name as product_category_name,
		COALESCE(r.numb_reviews, 0) AS numb_reviews,
		COALESCE(r.rating, 0) AS rating
	FROM products p
	JOIN product_product_categories ppc ON ppc.product_id = p.product_id
	JOIN product_categories pc ON pc.product_category_id = ppc.product_category_id
	LEFT JOIN (
		SELECT 
			product_id,
			COUNT(review_id) AS numb_reviews, 
			AVG(rating)  AS rating
		FROM reviews 
		GROUP BY product_id
	) r ON r.product_id = p.product_id
	ORDER BY %s
	LIMIT $1`, sort)

	rows, err := r.db.Query(sqlGetProduct, criteria.Limit)

	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get all product from database: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	products := make([]domain.ProductList, 0)
	for rows.Next() {
		var product domain.ProductList
		if err := rows.Scan(&product.ProductID, &product.Name, &product.Sku, &product.Brand, &product.Image, &product.Price, &product.ProductCategoryID,
			&product.ProductCategoryName, &product.NumbReviews, &product.Rating); err != nil {
			logger.Error("Error while scanning product category from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		products = append(products, product)
	}

	return products, nil
}

func (r ProductReposotory) GetAllPaginate(criteria *domain.ProductListCriteria) ([]domain.ProductList, int64, *errs.AppError) {
	var totalData int64
	var offset int64
	if criteria.Page > 0 {
		offset = (criteria.Page - 1) * criteria.Limit
	}

	searchName := fmt.Sprintf("%%%s%%", criteria.Keyword)

	sqlCountProduct := `
	SELECT 
		COUNT(p.product_id)
	FROM products p
	WHERE p.name ILIKE $1`

	err := r.db.QueryRow(sqlCountProduct, searchName).Scan(&totalData)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while count all product from database: " + err.Error())
		return nil, 0, errs.NewUnexpectedError("Unexpected database error")
	}

	sqlGetProduct := `
	SELECT 
		p.product_id, 
		p.name, 
		p.sku, 
		p.brand, 
		p.image, 
		p.price, 
		COALESCE(r.numb_reviews, 0) AS numb_reviews,
		COALESCE(r.rating, 0) AS rating
	FROM products p
	LEFT JOIN (
		SELECT 
			product_id,
			COUNT(review_id) AS numb_reviews, 
			AVG(rating)  AS rating
		FROM reviews 
		GROUP BY product_id
	) r ON r.product_id = p.product_id
	WHERE p.name ILIKE $1
	ORDER BY p.product_id ASC
	LIMIT $2
	OFFSET $3`

	rows, err := r.db.Query(sqlGetProduct, searchName, criteria.Limit, offset)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error while get all product from database: " + err.Error())
		return nil, 0, errs.NewUnexpectedError("Unexpected database error")
	}

	defer rows.Close()

	products := make([]domain.ProductList, 0)

	for rows.Next() {
		var product domain.ProductList
		if err := rows.Scan(&product.ProductID, &product.Name, &product.Sku, &product.Brand, &product.Image, &product.Price, &product.NumbReviews, &product.Rating); err != nil {
			logger.Error("Error while scanning get product from database: " + err.Error())
			return nil, 0, errs.NewUnexpectedError("Unexpected database error")
		}
		products = append(products, product)
	}

	return products, totalData, nil
}

func (r ProductReposotory) GetOneByID(productID int64) (*domain.ProductDetail, *errs.AppError) {

	var product domain.ProductDetail

	sqlGetProduct := `
	SELECT 
		p.product_id, 
		p.name, 
		p.sku, 
		p.brand, 
		p.image, 
		p.price, 
		p.description,
		p.stock
	FROM products p
	WHERE p.product_id = $1
	LIMIT 1`

	err := r.db.QueryRow(sqlGetProduct, productID).Scan(&product.ProductID, &product.Name, &product.Sku, &product.Brand, &product.Image, &product.Price,
		&product.Description, &product.Stock)
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

func (r ProductReposotory) Update(productID int64, data *domain.Product) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting update product: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlUpdate := `
	UPDATE products 
	SET name = $2, 
		sku = $3,
		brand = $4,
		image = $5,
		price = $6,
		description = $7,
		stock = $8,
		updated_at = $9
	WHERE product_id = $1`

	_, err = tx.Exec(sqlUpdate, productID, data.Name, data.Sku, data.Brand, data.Image, data.Price, data.Description, data.Stock, data.UpdatedAt)
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

func (r ProductReposotory) UpdateStock(productID, quantity int64) *errs.AppError {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting update stock: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	sqlUpdate := `
	UPDATE products 
	SET stock = stock - $2,
		updated_at = $3
	WHERE product_id = $1`

	_, err = tx.Exec(sqlUpdate, productID, quantity, time.Now())
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

func (r ProductReposotory) Delete(productID int64) *errs.AppError {

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

func getSortProduct(sort string, order string) string {
	var sortResult, orderResult string
	sort = strings.ToLower(sort)
	order = strings.ToLower(order)

	switch sort {
	case "numb_reviews":
		sortResult = "numb_reviews"
	default:
		sortResult = "product_id"
	}

	switch order {
	case "desc":
		orderResult = "desc"
	default:
		orderResult = "asc"
	}

	return fmt.Sprintf("%s %s", sortResult, orderResult)
}

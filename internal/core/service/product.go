package service

import (
	"fmt"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
)

type ProductService struct {
	repo                       port.ProductRepo
	productCategoryRepo        port.ProductCategoryRepo
	productProductCategoryRepo port.ProductProductCategoryRepo
	reviewRepo                 port.ReviewRepo
}

func NewProductService(repo port.ProductRepo, productCategoryRepo port.ProductCategoryRepo, productProductCategoryRepo port.ProductProductCategoryRepo, reviewRepo port.ReviewRepo) port.ProductService {
	return &ProductService{
		repo:                       repo,
		productCategoryRepo:        productCategoryRepo,
		productProductCategoryRepo: productProductCategoryRepo,
		reviewRepo:                 reviewRepo,
	}
}

func (r ProductService) Create(form *domain.Product) *errs.AppError {

	checkProduct, appErr := r.repo.CheckBySKU(form.Sku)
	if appErr != nil {
		return appErr
	}

	if checkProduct {
		errorMessage := fmt.Sprintf("SKU %s is already used", form.Sku)
		return errs.NewBadRequestError(errorMessage)
	}

	for _, productCategoryID := range form.ProductCategoryIDs {
		checkProductCategory, appErr := r.productCategoryRepo.CheckByID(productCategoryID)
		if appErr != nil {
			return appErr
		}

		if !checkProductCategory {
			return errs.NewBadRequestError("Product category not found")
		}
	}

	form.CreatedAt = time.Now().Format(dbTSLayout)
	form.UpdatedAt = time.Now().Format(dbTSLayout)

	newProductData, appErr := r.repo.Insert(form)
	if appErr != nil {
		return appErr
	}

	formProductProductCategory := make([]domain.ProductProductCategory, 0)

	for _, productCategoryID := range form.ProductCategoryIDs {
		formProductProductCategory = append(formProductProductCategory, domain.ProductProductCategory{
			ProductID:         newProductData.ProductID,
			ProductCategoryID: productCategoryID,
		})
	}

	appErr = r.productProductCategoryRepo.BulkInsert(formProductProductCategory)
	if appErr != nil {
		return appErr
	}

	return nil
}

func (r ProductService) GetList(criteria *domain.ProductListCriteria) ([]domain.ProductDetail, *errs.AppError) {

	products, appErr := r.repo.GetAll(criteria)
	if appErr != nil {
		return nil, appErr
	}
	result := make([]domain.ProductDetail, 0)
	for _, value := range products {
		var product domain.ProductDetail
		product.ProductID = value.ProductID
		product.Name = value.Name
		product.Sku = value.Sku
		product.Image = value.Image
		product.Brand = value.Brand
		product.Price = value.Price
		product.NumbReviews = value.NumbReviews
		product.Rating = value.Rating
		result = append(result, product)
	}

	return result, nil
}

func (r ProductService) GetListPaginate(criteria *domain.ProductListCriteria) ([]domain.ProductDetail, int64, *errs.AppError) {

	products, total, appErr := r.repo.GetAllPaginate(criteria)
	if appErr != nil {
		return nil, 0, appErr
	}

	result := make([]domain.ProductDetail, 0)
	for _, value := range products {
		var product domain.ProductDetail
		product.ProductID = value.ProductID
		product.Name = value.Name
		product.Sku = value.Sku
		product.Image = value.Image
		product.Brand = value.Brand
		product.Price = value.Price
		product.NumbReviews = value.NumbReviews
		product.Rating = value.Rating

		productCategories, appErr := r.productCategoryRepo.GetAllByProductID(value.ProductID)
		if appErr != nil {
			return nil, 0, appErr
		}

		product.ProductCategories = productCategories
		result = append(result, product)
	}

	return result, total, nil
}

func (r ProductService) GetDetail(productID int64) (*domain.ProductDetail, *errs.AppError) {
	// fetch product
	product, err := r.repo.GetOneByID(productID)
	if err != nil {
		return nil, err
	}

	// fetch product categories
	productCategories, err := r.productCategoryRepo.GetAllByProductID(productID)
	if err != nil {
		return nil, err
	}
	product.ProductCategories = productCategories

	// fetch product reviews
	productReviews, err := r.reviewRepo.GetAllByProductID(productID)
	if err != nil {
		return nil, err
	}
	product.Review = productReviews

	// calculate rating and number of reviews
	var totalRating int64
	for _, review := range productReviews {
		totalRating += int64(review.Rating)
	}

	if len(productReviews) > 0 {
		product.Rating = float32(totalRating) / float32(len(productReviews))
	}
	product.NumbReviews = int64(len(productReviews))

	return product, nil
}

func (r ProductService) Update(productID int64, form *domain.Product) *errs.AppError {

	checkProduct, appErr := r.repo.CheckByID(productID)
	if appErr != nil {
		return appErr
	}

	if !checkProduct {
		return errs.NewBadRequestError("Product not found")
	}

	checkProductSKU, appErr := r.repo.CheckByIDAndSKU(productID, form.Sku)
	if appErr != nil {
		return appErr
	}

	if checkProductSKU {
		errorMessage := fmt.Sprintf("SKU %s is already used", form.Sku)
		return errs.NewBadRequestError(errorMessage)
	}

	for _, productCategoryID := range form.ProductCategoryIDs {
		checkProductCategory, appErr := r.productCategoryRepo.CheckByID(productCategoryID)
		if appErr != nil {
			return appErr
		}

		if !checkProductCategory {
			return errs.NewBadRequestError("Product category not found")
		}
	}

	form.UpdatedAt = time.Now().Format(dbTSLayout)
	appErr = r.repo.Update(productID, form)
	if appErr != nil {
		return appErr
	}

	formProductProductCategory := make([]domain.ProductProductCategory, 0)

	for _, productCategoryID := range form.ProductCategoryIDs {
		formProductProductCategory = append(formProductProductCategory, domain.ProductProductCategory{
			ProductID:         productID,
			ProductCategoryID: productCategoryID,
		})
	}

	appErr = r.productProductCategoryRepo.DeleteAll(productID)
	if appErr != nil {
		return appErr
	}

	appErr = r.productProductCategoryRepo.BulkInsert(formProductProductCategory)
	if appErr != nil {
		return appErr
	}

	return nil
}

func (r ProductService) Delete(productID int64) *errs.AppError {

	checkProduct, appErr := r.repo.CheckByID(productID)
	if appErr != nil {
		return appErr
	}
	if !checkProduct {
		return errs.NewBadRequestError("Product not found")
	}

	appErr = r.repo.Delete(productID)
	if appErr != nil {
		return appErr
	}

	appErr = r.productProductCategoryRepo.DeleteAll(productID)
	if appErr != nil {
		return appErr
	}

	return nil
}

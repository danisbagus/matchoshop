package service

import (
	"fmt"
	"sync"
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

	var product *domain.ProductDetail
	var productCategories []domain.ProductCategory
	var productReviews []domain.Review

	productChan := make(chan *domain.ProductDetail, 1)
	productCategoriesChan := make(chan []domain.ProductCategory, 1)
	productReviewsChan := make(chan []domain.Review, 1)
	errorChan := make(chan *errs.AppError, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	// get detail product
	go func() {
		defer wg.Done()

		product, appErr := r.repo.GetOneByID(productID)
		if appErr != nil {
			errorChan <- appErr
		}

		productChan <- product
	}()

	// get detail product category
	go func() {
		defer wg.Done()

		productCategories, appErr := r.productCategoryRepo.GetAllByProductID(productID)
		if appErr != nil {
			errorChan <- appErr
		}

		productCategoriesChan <- productCategories

	}()

	// get list review
	go func() {
		defer wg.Done()

		reviews, appErr := r.reviewRepo.GetAllByProductID(productID)
		if appErr != nil {
			errorChan <- appErr
		}

		productReviewsChan <- reviews

	}()

	wg.Wait()

	close(errorChan)
	close(productChan)
	close(productCategoriesChan)
	close(productReviewsChan)

	for appErr := range errorChan {
		if appErr != nil {
			return nil, appErr
		}
	}

	for dataChan := range productChan {
		product = dataChan
	}

	for dataChan := range productCategoriesChan {
		productCategories = dataChan
	}

	product.ProductCategories = productCategories

	for dataChan := range productReviewsChan {
		productReviews = dataChan
	}

	product.Review = productReviews

	var totalRating int64
	var averageRating float32
	for _, value := range productReviews {
		totalRating += int64(value.Rating)
	}

	if totalRating > 0 {
		averageRating = (float32(totalRating)) / (float32(len(productReviews)))
	}
	product.Rating = averageRating
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

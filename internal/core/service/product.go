package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/danisbagus/matchoshop/internal/repository"
)

type ProductService struct {
	productRepo                repository.IProductRepository
	productCategoryRepo        repository.IProductCategoryRepository
	productProductCategoryRepo repository.IProductProductCategoryRepository
	reviewRepo                 repository.IReviewRepository
}

func NewProductService(repository repository.RepositoryCollection) port.ProductService {
	return &ProductService{
		productRepo:                repository.ProductReposotory,
		productCategoryRepo:        repository.ProductCategoryRepository,
		productProductCategoryRepo: repository.ProductProductCategoryRepository,
		reviewRepo:                 repository.ReviewRepository,
	}
}

func (r ProductService) Create(form *domain.Product) *errs.AppError {

	checkProduct, appErr := r.productRepo.CheckBySKU(form.Sku)
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

	newProductData, appErr := r.productRepo.Insert(form)
	if appErr != nil {
		return appErr
	}

	formProductProductCategory := make([]domain.ProductProductCategoryModel, 0)

	for _, productCategoryID := range form.ProductCategoryIDs {
		formProductProductCategory = append(formProductProductCategory, domain.ProductProductCategoryModel{
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

	products, appErr := r.productRepo.GetAll(criteria)
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

	products, total, appErr := r.productRepo.GetAllPaginate(criteria)
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
	var productCategories []domain.ProductCategoryModel
	var productReviews []domain.Review

	productChan := make(chan *domain.ProductDetail, 1)
	productCategoriesChan := make(chan []domain.ProductCategoryModel, 1)
	productReviewsChan := make(chan []domain.Review, 1)
	errorChan := make(chan *errs.AppError, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	// get detail product
	go func() {
		defer wg.Done()

		product, appErr := r.productRepo.GetOneByID(productID)
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

	checkProduct, appErr := r.productRepo.CheckByID(productID)
	if appErr != nil {
		return appErr
	}

	if !checkProduct {
		return errs.NewBadRequestError("Product not found")
	}

	checkProductSKU, appErr := r.productRepo.CheckByIDAndSKU(productID, form.Sku)
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
	appErr = r.productRepo.Update(productID, form)
	if appErr != nil {
		return appErr
	}

	formProductProductCategory := make([]domain.ProductProductCategoryModel, 0)

	for _, productCategoryID := range form.ProductCategoryIDs {
		formProductProductCategory = append(formProductProductCategory, domain.ProductProductCategoryModel{
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

	checkProduct, appErr := r.productRepo.CheckByID(productID)
	if appErr != nil {
		return appErr
	}
	if !checkProduct {
		return errs.NewBadRequestError("Product not found")
	}

	appErr = r.productRepo.Delete(productID)
	if appErr != nil {
		return appErr
	}

	appErr = r.productProductCategoryRepo.DeleteAll(productID)
	if appErr != nil {
		return appErr
	}

	return nil
}

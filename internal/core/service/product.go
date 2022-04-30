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
}

func NewProductService(repo port.ProductRepo, productCategoryRepo port.ProductCategoryRepo, productProductCategoryRepo port.ProductProductCategoryRepo) port.ProductService {
	return &ProductService{
		repo:                       repo,
		productCategoryRepo:        productCategoryRepo,
		productProductCategoryRepo: productProductCategoryRepo,
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

func (r ProductService) GetList() ([]domain.ProductDetail, *errs.AppError) {

	products, appErr := r.repo.GetAll()
	if appErr != nil {
		return nil, appErr
	}

	result := make([]domain.ProductDetail, 0)
	mapProduct := make(map[int64]domain.ProductDetail)

	for _, value := range products {
		var productCategory domain.ProductCategory
		productCategory.ProductCategoryID = value.ProductCategoryID
		productCategory.Name = value.ProductCategoryName

		if mapValue, ok := mapProduct[value.ProductID]; ok {
			mapValue.ProductCategories = append(mapValue.ProductCategories, productCategory)
			mapProduct[value.ProductID] = mapValue
		} else {
			var product domain.ProductDetail
			product.ProductID = value.ProductID
			product.Name = value.Name
			product.Sku = value.Sku
			product.Image = value.Image
			product.Brand = value.Brand
			product.Price = value.Price
			product.ProductCategories = append(product.ProductCategories, productCategory)
			mapProduct[value.ProductID] = product
		}
	}

	for _, valData := range mapProduct {
		result = append(result, valData)
	}

	return result, nil
}

func (r ProductService) GetDetail(productID int64) (*domain.ProductDetail, *errs.AppError) {

	var product *domain.ProductDetail
	var productCategories []domain.ProductCategory

	productChan := make(chan *domain.ProductDetail, 1)
	productCategoriesChan := make(chan []domain.ProductCategory, 1)
	errorChan := make(chan *errs.AppError, 2)

	var wg sync.WaitGroup
	wg.Add(2)

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

	wg.Wait()

	close(errorChan)
	close(productChan)
	close(productCategoriesChan)

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

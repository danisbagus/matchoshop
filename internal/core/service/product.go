package service

import (
	"fmt"
	"sync"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type ProductService struct {
	repo                       port.IProductRepo
	productCategoryRepo        port.IProductCategoryRepo
	productProductCategoryRepo port.IProductProductCategoryRepo
}

func NewProductService(repo port.IProductRepo, productCategoryRepo port.IProductCategoryRepo, productProductCategoryRepo port.IProductProductCategoryRepo) port.IProductService {
	return &ProductService{
		repo:                       repo,
		productCategoryRepo:        productCategoryRepo,
		productProductCategoryRepo: productProductCategoryRepo,
	}
}

func (r ProductService) Create(req *dto.CreateProductRequest) (*dto.ResponseData, *errs.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	checkProduct, appErr := r.repo.CheckBySKU(req.Sku)
	if appErr != nil {
		return nil, appErr
	}

	if checkProduct {
		errorMessage := fmt.Sprintf("SKU %s is already used", req.Sku)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	for _, productCategoryID := range req.ProductCategoryID {
		checkProductCategory, appErr := r.productCategoryRepo.CheckByID(productCategoryID)
		if appErr != nil {
			return nil, appErr
		}

		if !checkProductCategory {
			return nil, errs.NewBadRequestError("Product category not found")
		}
	}

	formProduct := domain.Product{
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now().Format(dbTSLayout),
		UpdatedAt:   time.Now().Format(dbTSLayout),
	}

	newProductData, appErr := r.repo.Insert(&formProduct)
	if appErr != nil {
		return nil, appErr
	}

	formProductProductCategory := make([]domain.ProductProductCategory, 0)

	for _, productCategoryID := range req.ProductCategoryID {
		formProductProductCategory = append(formProductProductCategory, domain.ProductProductCategory{
			ProductID:         newProductData.ProductID,
			ProductCategoryID: productCategoryID,
		})
	}

	appErr = r.productProductCategoryRepo.BulkInsert(formProductProductCategory)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewCreateProductResponse("Sucessfully create data", newProductData)
	return response, nil
}

func (r ProductService) GetList() (*dto.ResponseData, *errs.AppError) {

	products, appErr := r.repo.GetAll()
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewGetProductListResponse("Successfully get data", products)

	return response, nil
}

func (r ProductService) GetDetail(productID int64) (*dto.ResponseData, *errs.AppError) {

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

	response := dto.NewGetProductDetailResponse("Successfully get data", product)

	return response, nil
}

func (r ProductService) Update(productID int64, req *dto.CreateProductRequest) (*dto.ResponseData, *errs.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	checkProduct, appErr := r.repo.CheckByID(productID)
	if appErr != nil {
		return nil, appErr
	}

	if !checkProduct {
		return nil, errs.NewBadRequestError("Product not found")
	}

	checkProductSKU, appErr := r.repo.CheckByIDAndSKU(productID, req.Sku)
	if appErr != nil {
		return nil, appErr
	}

	if checkProductSKU {
		errorMessage := fmt.Sprintf("SKU %s is already used", req.Sku)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	for _, productCategoryID := range req.ProductCategoryID {
		checkProductCategory, appErr := r.productCategoryRepo.CheckByID(productCategoryID)
		if appErr != nil {
			return nil, appErr
		}

		if !checkProductCategory {
			return nil, errs.NewBadRequestError("Product category not found")
		}
	}

	formProduct := domain.Product{
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now().Format(dbTSLayout),
		UpdatedAt:   time.Now().Format(dbTSLayout),
	}

	appErr = r.repo.Update(productID, &formProduct)
	if appErr != nil {
		return nil, appErr
	}

	formProductProductCategory := make([]domain.ProductProductCategory, 0)

	for _, productCategoryID := range req.ProductCategoryID {
		formProductProductCategory = append(formProductProductCategory, domain.ProductProductCategory{
			ProductID:         productID,
			ProductCategoryID: productCategoryID,
		})
	}

	appErr = r.productProductCategoryRepo.DeleteAll(productID)
	if appErr != nil {
		return nil, appErr
	}

	appErr = r.productProductCategoryRepo.BulkInsert(formProductProductCategory)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.GenerateResponseData("Successfully update data", map[string]string{})

	return response, nil
}

func (r ProductService) Delete(productID int64) (*dto.ResponseData, *errs.AppError) {

	checkProduct, appErr := r.repo.CheckByID(productID)
	if appErr != nil {
		return nil, appErr
	}

	if !checkProduct {
		return nil, errs.NewBadRequestError("Product not found")
	}

	appErr = r.repo.Delete(productID)
	if appErr != nil {
		return nil, appErr
	}

	appErr = r.productProductCategoryRepo.DeleteAll(productID)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.GenerateResponseData("Successfully delete data", map[string]string{})

	return response, nil
}

package service

import (
	"fmt"
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

	checkProduct, appErr := r.repo.CheckBySKUAndMerchantID(req.Sku, req.MerchantID)
	if appErr != nil {
		return nil, appErr
	}

	if checkProduct {
		errorMessage := fmt.Sprintf("SKU %s is already used", req.Sku)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	for _, productCategoryID := range req.ProductCategoryID {
		checkProductCategory, appErr := r.productCategoryRepo.CheckByIDAndMerchantID(productCategoryID, req.MerchantID)
		if appErr != nil {
			return nil, appErr
		}

		if !checkProductCategory {
			return nil, errs.NewBadRequestError("Product category not found")
		}
	}

	formProduct := domain.Product{
		MerchantID:  req.MerchantID,
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

func (r ProductService) GetList(merchantID int64) (*dto.ResponseData, *errs.AppError) {

	products, appErr := r.repo.GetAllByMerchantID(merchantID)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewGetProductListResponse("Successfully get data", products)

	return response, nil
}

func (r ProductService) GetDetail(productID int64, merchantID int64) (*dto.ResponseData, *errs.AppError) {

	product, appErr := r.repo.GetOneByIDAndMerchantID(productID, merchantID)
	if appErr != nil {
		return nil, appErr
	}

	productCategories, appErr := r.productCategoryRepo.GetAllByProductIDAndMerchantID(productID, merchantID)
	if appErr != nil {
		return nil, appErr
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

	checkProduct, appErr := r.repo.CheckByIDAndMerchantID(productID, req.MerchantID)
	if appErr != nil {
		return nil, appErr
	}

	if !checkProduct {
		return nil, errs.NewBadRequestError("Product not found")
	}

	checkProductSKU, appErr := r.repo.CheckByIDAndSKUAndMerchantID(productID, req.Sku, req.MerchantID)
	if appErr != nil {
		return nil, appErr
	}

	if checkProductSKU {
		errorMessage := fmt.Sprintf("SKU %s is already used", req.Sku)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	for _, productCategoryID := range req.ProductCategoryID {
		checkProductCategory, appErr := r.productCategoryRepo.CheckByIDAndMerchantID(productCategoryID, req.MerchantID)
		if appErr != nil {
			return nil, appErr
		}

		if !checkProductCategory {
			return nil, errs.NewBadRequestError("Product category not found")
		}
	}

	formProduct := domain.Product{
		MerchantID:  req.MerchantID,
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

func (r ProductService) Delete(productID int64, merchantID int64) (*dto.ResponseData, *errs.AppError) {

	checkProduct, appErr := r.repo.CheckByIDAndMerchantID(productID, merchantID)
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

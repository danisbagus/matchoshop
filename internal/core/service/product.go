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

	checkProductCategory, appErr := r.productCategoryRepo.CheckByIDAndMerchantID(req.ProductCategoryID, req.MerchantID)
	if appErr != nil {
		return nil, appErr
	}

	if !checkProductCategory {
		return nil, errs.NewBadRequestError("Product category not found")
	}

	if checkProduct {
		errorMessage := fmt.Sprintf("Product with SKU %s is already exits", req.Sku)
		return nil, errs.NewBadRequestError(errorMessage)
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

	formProductProductCategory := domain.ProductProductCategory{
		ProductID:         newProductData.ProductID,
		ProductCategoryID: req.ProductCategoryID,
	}

	appErr = r.productProductCategoryRepo.Insert(&formProductProductCategory)
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

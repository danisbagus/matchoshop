package service

import (
	"testing"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockProductRepo = &mocks.IProductRepo{Mock: mock.Mock{}}
var mockProductProductRepo = &mocks.IProductProductCategoryRepo{Mock: mock.Mock{}}
var productService = ProductService{repo: mockProductRepo, productCategoryRepo: mockProductCategoryRepo, productProductCategoryRepo: mockProductProductRepo}

func TestProduct_Create_NotValidated(t *testing.T) {

	req := dto.CreateProductRequest{}

	product, appErr := productService.Create(&req)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)
}

func TestProduct_Create_SKUExits(t *testing.T) {

	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 15 Modern",
		Sku:               "SKU001",
		Description:       "The modern TB",
		ProductCategoryID: []int64{1},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckBySKU", req.Sku).Return(true, nil)

	product, appErr := productService.Create(&req)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)
}

func TestProduct_Create_ProductCategoryNotFound(t *testing.T) {

	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 15 Modern",
		Sku:               "SKU001",
		Description:       "The modern TB",
		ProductCategoryID: []int64{1},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckBySKU", req.Sku).Return(false, nil)

	mockProductCategoryRepo.Mock.On("CheckByID", req.ProductCategoryID[0]).Return(false, nil)

	product, appErr := productService.Create(&req)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)
	// assert.Equal(t, appErr.Message, "Product category not found")
}

func TestProduct_Create_FailedInsertProduct(t *testing.T) {

	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 15 Modern",
		Sku:               "SKU0012345678901234567890",
		Description:       "The modern TB",
		ProductCategoryID: []int64{2},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckBySKU", req.Sku).Return(false, nil)

	mockProductCategoryRepo.Mock.On("CheckByID", req.ProductCategoryID[0]).Return(true, nil)

	formProduct := domain.Product{
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now().Format(dbTSLayout),
		UpdatedAt:   time.Now().Format(dbTSLayout),
	}

	mockProductRepo.Mock.On("Insert", &formProduct).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	product, appErr := productService.Create(&req)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)
}

func TestProduct_Create_FailedInsertProductProductCategory(t *testing.T) {

	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 15 Modern",
		Sku:               "SKU001234",
		Description:       "The modern TB",
		ProductCategoryID: []int64{2},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckBySKU", req.Sku).Return(false, nil)

	mockProductCategoryRepo.Mock.On("CheckByID", req.ProductCategoryID[0]).Return(true, nil)

	formProduct := domain.Product{
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now().Format(dbTSLayout),
		UpdatedAt:   time.Now().Format(dbTSLayout),
	}

	returnInsertProduct := domain.Product{
		ProductID:   0,
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
	}

	mockProductRepo.Mock.On("Insert", &formProduct).Return(&returnInsertProduct, nil)

	formProductProductCategory := []domain.ProductProductCategory{
		{ProductID: returnInsertProduct.ProductID, ProductCategoryID: req.ProductCategoryID[0]},
	}

	mockProductProductRepo.Mock.On("BulkInsert", formProductProductCategory).Return(errs.NewUnexpectedError("Unexpected database error"))

	product, appErr := productService.Create(&req)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)

}

func TestProduct_Create_Success(t *testing.T) {

	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 16 Modern",
		Sku:               "SKU001234",
		Description:       "The modern TB",
		ProductCategoryID: []int64{2},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckBySKU", req.Sku).Return(false, nil)

	mockProductCategoryRepo.Mock.On("CheckByID", req.ProductCategoryID[0]).Return(true, nil)

	formProduct := domain.Product{
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now().Format(dbTSLayout),
		UpdatedAt:   time.Now().Format(dbTSLayout),
	}

	returnInsertProduct := domain.Product{
		ProductID:   1,
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
	}

	mockProductRepo.Mock.On("Insert", &formProduct).Return(&returnInsertProduct, nil)

	formProductProductCategory := []domain.ProductProductCategory{
		{ProductID: returnInsertProduct.ProductID, ProductCategoryID: req.ProductCategoryID[0]},
	}

	mockProductProductRepo.Mock.On("BulkInsert", formProductProductCategory).Return(nil)

	product, appErr := productService.Create(&req)

	assert.NotNil(t, product)
	assert.Nil(t, appErr)
}

func TestProduct_GetDetail_NotFound(t *testing.T) {

	var productID int64 = 10

	mockProductRepo.Mock.On("GetOneByID", productID).Return(nil, errs.NewNotFoundError("Product not found!"))

	product, appErr := productService.GetDetail(productID)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)
}

func TestProduct_GetDetail_Success(t *testing.T) {

	var productID int64 = 1

	productDetailResult := domain.ProductDetail{
		Name:        "TV Elkoga 16 Modern",
		Sku:         "SKU001234",
		Description: "The modern TB",
		Price:       10000,
	}

	mockProductRepo.Mock.On("GetOneByID", productID).Return(&productDetailResult, nil)

	ProductCategoriesResult := []domain.ProductCategory{
		{
			ProductCategoryID: 1,
			Name:              "Modern shoes",
		},
	}

	mockProductCategoryRepo.Mock.On("GetAllByProductID", productID).Return(ProductCategoriesResult, nil)

	product, appErr := productService.GetDetail(productID)

	assert.NotNil(t, product)
	assert.Nil(t, appErr)
}

func TestProduct_Update_NotValidated(t *testing.T) {

	var productID int64 = 10
	req := dto.CreateProductRequest{}

	product, appErr := productService.Update(productID, &req)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)
}

func TestProduct_Update_NotFound(t *testing.T) {

	var productID int64 = 10
	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 16 Modern",
		Sku:               "SKU001234",
		Description:       "The modern TB",
		ProductCategoryID: []int64{2},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckByID", productID).Return(false, nil)

	product, appErr := productService.Update(productID, &req)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)
}

func TestProduct_Update_SKUExits(t *testing.T) {

	var productID int64 = 1
	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 16 Modern",
		Sku:               "SKU00123",
		Description:       "The modern TB",
		ProductCategoryID: []int64{2},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckByID", productID).Return(true, nil)

	mockProductRepo.Mock.On("CheckByIDAndSKU", productID, req.Sku).Return(true, nil)

	product, appErr := productService.Update(productID, &req)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)
}

func TestProduct_Update_ProductCategoryNotFound(t *testing.T) {

	var productID int64 = 1
	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 16 Modern",
		Sku:               "SKU0012345",
		Description:       "The modern TB",
		ProductCategoryID: []int64{1},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckByID", productID).Return(true, nil)

	mockProductRepo.Mock.On("CheckByIDAndSKU", productID, req.Sku).Return(false, nil)

	mockProductCategoryRepo.Mock.On("CheckByID", req.ProductCategoryID[0]).Return(false, nil)

	product, appErr := productService.Update(productID, &req)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)

}

func TestProduct_Update_FailedUpdateProduct(t *testing.T) {

	var productID int64 = 1
	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 16 Modern",
		Sku:               "SKU0012345678901234567890",
		Description:       "The modern TB",
		ProductCategoryID: []int64{2},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckByID", productID).Return(true, nil)

	mockProductRepo.Mock.On("CheckByIDAndSKU", productID, req.Sku).Return(false, nil)

	mockProductCategoryRepo.Mock.On("CheckByID", req.ProductCategoryID[0]).Return(true, nil)

	formProduct := domain.Product{
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now().Format(dbTSLayout),
		UpdatedAt:   time.Now().Format(dbTSLayout),
	}

	mockProductRepo.Mock.On("Update", productID, &formProduct).Return(errs.NewUnexpectedError("Unexpected database error"))

	product, appErr := productService.Update(productID, &req)

	// fmt.Println(appErr)

	assert.Nil(t, product)
	assert.NotNil(t, appErr)
}

func TestProduct_Update_Success(t *testing.T) {

	var productID int64 = 1
	req := dto.CreateProductRequest{
		Name:              "TV Elkoga 16 Modern",
		Sku:               "SKU0012345",
		Description:       "The modern TB",
		ProductCategoryID: []int64{2},
		Price:             10000,
	}

	mockProductRepo.Mock.On("CheckByID", productID).Return(true, nil)

	mockProductRepo.Mock.On("CheckByIDAndSKU", productID, req.Sku).Return(false, nil)

	mockProductCategoryRepo.Mock.On("CheckByID", req.ProductCategoryID[0]).Return(true, nil)

	formProduct := domain.Product{
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now().Format(dbTSLayout),
		UpdatedAt:   time.Now().Format(dbTSLayout),
	}

	mockProductRepo.Mock.On("Update", productID, &formProduct).Return(nil)

	mockProductProductRepo.Mock.On("DeleteAll", productID).Return(nil)

	formProductProductCategory := []domain.ProductProductCategory{
		{ProductID: productID, ProductCategoryID: req.ProductCategoryID[0]},
	}

	mockProductProductRepo.Mock.On("BulkInsert", formProductProductCategory).Return(nil)

	product, appErr := productService.Update(productID, &req)

	assert.NotNil(t, product)
	assert.Nil(t, appErr)
}
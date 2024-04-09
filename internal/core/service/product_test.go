package service

import (
	"testing"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/repository"
	"github.com/danisbagus/matchoshop/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	description = "The modern TB"
)

func setupProductTest(t *testing.T) (mocks.RepoCollectionMocks, port.ProductService) {
	repoMock := mocks.RepoCollectionMocks{
		ProductReposotory:                mocks.NewIProductRepository(t),
		ProductCategoryRepository:        mocks.NewIProductCategoryRepository(t),
		ProductProductCategoryRepository: mocks.NewIProductProductCategoryRepository(t),
	}

	repoCollection := repository.RepositoryCollection{
		ProductReposotory:                repoMock.ProductReposotory,
		ProductProductCategoryRepository: repoMock.ProductProductCategoryRepository,
	}

	service := NewProductService(repoCollection)
	return repoMock, service
}

func TestProduct_Create_NotValidated(t *testing.T) {
	_, service := setupProductTest(t)

	form := new(domain.Product)

	appErr := service.Create(form)
	assert.NotNil(t, appErr)
}

func TestProduct_Create_SKUExits(t *testing.T) {
	repoMock, service := setupProductTest(t)

	form := &domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 15 Modern",
			Sku:         "SKU001",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{1},
	}

	repoMock.ProductReposotory.Mock.On("CheckBySKU", form.Sku).Return(true, nil)

	appErr := service.Create(form)
	assert.NotNil(t, appErr)
}

func TestProduct_Create_ProductCategoryNotFound(t *testing.T) {
	repoMock, service := setupProductTest(t)

	form := &domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 15 Modern",
			Sku:         "SKU001",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{1},
	}

	repoMock.ProductReposotory.Mock.On("CheckBySKU", form.Sku).Return(false, nil)
	repoMock.ProductReposotory.Mock.On("CheckByID", form.ProductCategoryIDs[0]).Return(false, nil)

	appErr := service.Create(form)
	assert.NotNil(t, appErr)
}

func TestProduct_Create_FailedInsertProduct(t *testing.T) {
	repoMock, service := setupProductTest(t)

	form := &domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 15 Modern",
			Sku:         "SKU0012345678901234567890",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{2},
	}

	repoMock.ProductReposotory.Mock.On("CheckBySKU", form.Sku).Return(false, nil)
	repoMock.ProductCategoryRepository.Mock.On("CheckByID", form.ProductCategoryIDs[0]).Return(true, nil)

	formProduct := domain.Product{
		ProductModel: domain.ProductModel{
			Name:        form.Name,
			Sku:         form.Sku,
			Description: form.Description,
			Price:       form.Price,
			CreatedAt:   time.Now().Format(dbTSLayout),
			UpdatedAt:   time.Now().Format(dbTSLayout),
		},
	}

	repoMock.ProductReposotory.Mock.On("Insert", &formProduct).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	appErr := service.Create(form)
	assert.NotNil(t, appErr)
}

func TestProduct_Create_FailedInsertProductProductCategory(t *testing.T) {
	repoMock, service := setupProductTest(t)

	form := &domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 15 Modern",
			Sku:         "SKU001234",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{2},
	}

	repoMock.ProductReposotory.Mock.On("CheckBySKU", form.Sku).Return(false, nil)

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", form.ProductCategoryIDs[0]).Return(true, nil)

	formProduct := domain.Product{
		ProductModel: domain.ProductModel{
			Name:        form.Name,
			Sku:         form.Sku,
			Description: form.Description,
			Price:       form.Price,
			CreatedAt:   time.Now().Format(dbTSLayout),
			UpdatedAt:   time.Now().Format(dbTSLayout),
		},
	}

	returnInsertProduct := domain.Product{
		ProductModel: domain.ProductModel{
			ProductID:   0,
			Name:        form.Name,
			Sku:         form.Sku,
			Description: form.Description,
			Price:       form.Price,
		},
	}

	repoMock.ProductReposotory.Mock.On("Insert", &formProduct).Return(&returnInsertProduct, nil)

	formProductProductCategory := []domain.ProductProductCategory{
		{ProductID: returnInsertProduct.ProductID, ProductCategoryID: form.ProductCategoryIDs[0]},
	}

	repoMock.ProductCategoryRepository.Mock.On("BulkInsert", formProductProductCategory).Return(errs.NewUnexpectedError("Unexpected database error"))

	appErr := service.Create(form)
	assert.NotNil(t, appErr)

}

func TestProduct_Create_Success(t *testing.T) {
	repoMock, service := setupProductTest(t)

	form := &domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 16 Modern",
			Sku:         "SKU001234",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{2},
	}

	repoMock.ProductReposotory.Mock.On("CheckBySKU", form.Sku).Return(false, nil)

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", form.ProductCategoryIDs[0]).Return(true, nil)

	formProduct := domain.Product{
		ProductModel: domain.ProductModel{
			Name:        form.Name,
			Sku:         form.Sku,
			Description: form.Description,
			Price:       form.Price,
			CreatedAt:   time.Now().Format(dbTSLayout),
			UpdatedAt:   time.Now().Format(dbTSLayout),
		},
	}

	returnInsertProduct := domain.Product{
		ProductModel: domain.ProductModel{
			ProductID:   1,
			Name:        form.Name,
			Sku:         form.Sku,
			Description: form.Description,
			Price:       form.Price,
		},
	}

	repoMock.ProductReposotory.Mock.On("Insert", &formProduct).Return(&returnInsertProduct, nil)

	formProductProductCategory := []domain.ProductProductCategory{
		{ProductID: returnInsertProduct.ProductID, ProductCategoryID: form.ProductCategoryIDs[0]},
	}

	repoMock.ProductCategoryRepository.Mock.On("BulkInsert", formProductProductCategory).Return(nil)

	appErr := service.Create(form)

	assert.Nil(t, appErr)
}

func TestProduct_GetDetail_NotFound(t *testing.T) {
	repoMock, service := setupProductTest(t)

	var productID int64 = 10

	repoMock.ProductReposotory.Mock.On("GetOneByID", productID).Return(nil, errs.NewNotFoundError("Product not found!"))

	ProductCategoriesResult := []domain.ProductCategory{
		{
			ProductCategoryID: 1,
			Name:              "Modern shoes",
		},
	}

	repoMock.ProductCategoryRepository.Mock.On("GetAllByProductID", productID).Return(ProductCategoriesResult, nil)

	product, appErr := service.GetDetail(productID)
	assert.Nil(t, product)
	assert.NotNil(t, appErr)
}

func TestProduct_GetDetail_Success(t *testing.T) {
	repoMock, service := setupProductTest(t)

	var productID int64 = 1
	productDetailResult := domain.ProductDetail{
		ProductModel: domain.ProductModel{

			Name:        "TV Elkoga 16 Modern",
			Sku:         "SKU001234",
			Description: &description,
			Price:       10000,
		},
	}

	repoMock.ProductReposotory.Mock.On("GetOneByID", productID).Return(&productDetailResult, nil)

	ProductCategoriesResult := []domain.ProductCategory{
		{
			ProductCategoryID: 1,
			Name:              "Modern shoes",
		},
	}

	repoMock.ProductCategoryRepository.Mock.On("GetAllByProductID", productID).Return(ProductCategoriesResult, nil)

	product, appErr := service.GetDetail(productID)
	assert.Nil(t, product)
	assert.Nil(t, appErr)
}

func TestProduct_Update_NotValidated(t *testing.T) {
	_, service := setupProductTest(t)

	var productID int64 = 10
	req := domain.Product{}

	appErr := service.Update(productID, &req)
	assert.NotNil(t, appErr)
}

func TestProduct_Update_NotFound(t *testing.T) {
	repoMock, service := setupProductTest(t)

	var productID int64 = 10
	req := domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 16 Modern",
			Sku:         "SKU001234",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{2},
	}

	repoMock.ProductReposotory.Mock.On("CheckByID", productID).Return(false, nil)

	appErr := service.Update(productID, &req)
	assert.NotNil(t, appErr)
}

func TestProduct_Update_SKUExits(t *testing.T) {
	repoMock, service := setupProductTest(t)

	var productID int64 = 1
	req := domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 16 Modern",
			Sku:         "SKU00123",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{2},
	}

	repoMock.ProductReposotory.Mock.On("CheckByID", productID).Return(true, nil)
	repoMock.ProductReposotory.Mock.On("CheckByIDAndSKU", productID, req.Sku).Return(true, nil)

	appErr := service.Update(productID, &req)
	assert.NotNil(t, appErr)
}

func TestProduct_Update_ProductCategoryNotFound(t *testing.T) {
	repoMock, service := setupProductTest(t)

	var productID int64 = 1
	req := domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 16 Modern",
			Sku:         "SKU0012345",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{1},
	}

	repoMock.ProductReposotory.Mock.On("CheckByID", productID).Return(true, nil)

	repoMock.ProductReposotory.Mock.On("CheckByIDAndSKU", productID, req.Sku).Return(false, nil)

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", req.ProductCategoryIDs[0]).Return(false, nil)

	appErr := service.Update(productID, &req)
	assert.NotNil(t, appErr)

}

func TestProduct_Update_FailedUpdateProduct(t *testing.T) {
	repoMock, service := setupProductTest(t)

	var productID int64 = 1
	req := domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 16 Modern",
			Sku:         "SKU0012345678901234567890",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{2},
	}

	repoMock.ProductReposotory.Mock.On("CheckByID", productID).Return(true, nil)

	repoMock.ProductReposotory.Mock.On("CheckByIDAndSKU", productID, req.Sku).Return(false, nil)

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", req.ProductCategoryIDs[0]).Return(true, nil)

	formProduct := domain.Product{
		ProductModel: domain.ProductModel{
			Name:        req.Name,
			Sku:         req.Sku,
			Description: req.Description,
			Price:       req.Price,
			CreatedAt:   time.Now().Format(dbTSLayout),
			UpdatedAt:   time.Now().Format(dbTSLayout),
		},
	}

	repoMock.ProductReposotory.Mock.On("Update", productID, &formProduct).Return(errs.NewUnexpectedError("Unexpected database error"))
	appErr := service.Update(productID, &req)

	assert.NotNil(t, appErr)
}

func TestProduct_Update_Success(t *testing.T) {
	repoMock, service := setupProductTest(t)

	var productID int64 = 1
	req := domain.Product{
		ProductModel: domain.ProductModel{
			Name:        "TV Elkoga 16 Modern",
			Sku:         "SKU0012345",
			Description: &description,
			Price:       10000,
		},
		ProductCategoryIDs: []int64{2},
	}

	repoMock.ProductReposotory.Mock.On("CheckByID", productID).Return(true, nil)

	repoMock.ProductReposotory.Mock.On("CheckByIDAndSKU", productID, req.Sku).Return(false, nil)

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", req.ProductCategoryIDs[0]).Return(true, nil)

	formProduct := domain.Product{
		ProductModel: domain.ProductModel{

			Name:        req.Name,
			Sku:         req.Sku,
			Description: req.Description,
			Price:       req.Price,
			CreatedAt:   time.Now().Format(dbTSLayout),
			UpdatedAt:   time.Now().Format(dbTSLayout),
		},
	}

	repoMock.ProductReposotory.Mock.On("Update", productID, &formProduct).Return(nil)

	repoMock.ProductCategoryRepository.Mock.On("DeleteAll", productID).Return(nil)

	formProductProductCategory := []domain.ProductProductCategory{
		{ProductID: productID, ProductCategoryID: req.ProductCategoryIDs[0]},
	}

	repoMock.ProductCategoryRepository.Mock.On("BulkInsert", formProductProductCategory).Return(nil)

	appErr := service.Update(productID, &req)
	assert.Nil(t, appErr)
}

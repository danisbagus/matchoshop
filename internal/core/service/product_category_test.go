package service

import (
	"testing"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/internal/repository"
	"github.com/danisbagus/matchoshop/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

func setupProductCategoryTest(t *testing.T) (mocks.RepoCollectionMocks, port.ProductCategoryService) {
	repoMock := mocks.RepoCollectionMocks{
		ProductCategoryRepository: mocks.NewIProductCategoryRepository(t),
	}

	repoCollection := repository.RepositoryCollection{
		ProductCategoryRepository: repoMock.ProductCategoryRepository,
	}

	service := NewProductCategoryService(repoCollection)
	return repoMock, service
}

func TestProductCategory_Create_NotValidated(t *testing.T) {
	_, service := setupProductCategoryTest(t)

	req := dto.CreateProductCategoryRequest{}
	productCategory, appErr := service.Create(&req)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Create_NameExits(t *testing.T) {
	repoMock, service := setupProductCategoryTest(t)

	req := dto.CreateProductCategoryRequest{
		Name: "Electronic",
	}

	repoMock.ProductCategoryRepository.Mock.On("CheckByName", req.Name).Return(true, nil)

	productCategory, appErr := service.Create(&req)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Create_Success(t *testing.T) {
	repoMock, service := setupProductCategoryTest(t)

	req := dto.CreateProductCategoryRequest{
		Name: "Sport",
	}

	repoMock.ProductCategoryRepository.Mock.On("CheckByName", req.Name).Return(false, nil)

	formProductCategory := domain.ProductCategory{
		Name:      req.Name,
		CreatedAt: time.Now().Format(dbTSLayout),
		UpdatedAt: time.Now().Format(dbTSLayout),
	}

	resultProductCategory := domain.ProductCategory{
		ProductCategoryID: 1,
		Name:              formProductCategory.Name,
		CreatedAt:         formProductCategory.CreatedAt,
		UpdatedAt:         formProductCategory.UpdatedAt,
	}
	repoMock.ProductCategoryRepository.Mock.On("Insert", &formProductCategory).Return(&resultProductCategory, nil)

	productCategory, appErr := service.Create(&req)

	assert.NotNil(t, productCategory)
	assert.Nil(t, appErr)
}
func TestProductCategory_GetDetail_NotFound(t *testing.T) {
	repoMock, service := setupProductCategoryTest(t)

	var productCategoryID int64 = 10

	repoMock.ProductCategoryRepository.Mock.On("GetOneByID", productCategoryID).Return(nil, errs.NewNotFoundError("Product category not found!"))

	productCategory, appErr := service.GetDetail(productCategoryID)
	assert.NotNil(t, appErr)
	assert.Nil(t, productCategory)
}

func TestProductCategory_GetDetail_Success(t *testing.T) {
	repoMock, service := setupProductCategoryTest(t)

	var productCategoryID int64 = 1

	productCategoryResult := domain.ProductCategory{
		ProductCategoryID: 1,
		Name:              "Modern shoes",
	}

	repoMock.ProductCategoryRepository.Mock.On("GetOneByID", productCategoryID).Return(&productCategoryResult, nil)

	productCategory, appErr := service.GetDetail(productCategoryID)
	assert.NotNil(t, productCategory)
	assert.Nil(t, appErr)
}

func TestProductCategory_Update_NotValidated(t *testing.T) {
	_, service := setupProductCategoryTest(t)

	productCategoryID := 1
	reqProductCategory := dto.CreateProductCategoryRequest{}

	productCategory, appErr := service.Update(int64(productCategoryID), &reqProductCategory)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Update_NotFound(t *testing.T) {
	repoMock, service := setupProductCategoryTest(t)

	productCategoryID := 1
	req := dto.CreateProductCategoryRequest{
		Name: "Electonics",
	}

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", int64(productCategoryID)).Return(false, nil)

	productCategory, appErr := service.Update(int64(productCategoryID), &req)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Update_NameExits(t *testing.T) {
	repoMock, service := setupProductCategoryTest(t)

	productCategoryID := 2
	req := dto.CreateProductCategoryRequest{
		Name: "Electonics",
	}

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", int64(productCategoryID)).Return(true, nil)
	repoMock.ProductCategoryRepository.Mock.On("CheckByIDAndName", int64(productCategoryID), req.Name).Return(true, nil)

	productCategory, appErr := service.Update(int64(productCategoryID), &req)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Updated_Success(t *testing.T) {
	repoMock, service := setupProductCategoryTest(t)

	productCategoryID := 2
	req := dto.CreateProductCategoryRequest{
		Name: "Sport",
	}

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", int64(productCategoryID)).Return(true, nil)
	repoMock.ProductCategoryRepository.Mock.On("CheckByIDAndName", int64(productCategoryID), req.Name).Return(false, nil)

	formProductCategory := domain.ProductCategory{
		Name:      req.Name,
		CreatedAt: time.Now().Format(dbTSLayout),
		UpdatedAt: time.Now().Format(dbTSLayout),
	}

	repoMock.ProductCategoryRepository.Mock.On("Update", int64(productCategoryID), &formProductCategory).Return(nil)

	productCategory, appErr := service.Update(int64(productCategoryID), &req)

	assert.NotNil(t, productCategory)
	assert.Nil(t, appErr)
}

func TestProductCategory_Delete_NotFound(t *testing.T) {
	repoMock, service := setupProductCategoryTest(t)

	productCategoryID := 1

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", int64(productCategoryID)).Return(false, nil)

	productCategory, appErr := service.Delete(int64(productCategoryID))

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Delete_Success(t *testing.T) {
	repoMock, service := setupProductCategoryTest(t)

	productCategoryID := 2

	repoMock.ProductCategoryRepository.Mock.On("CheckByID", int64(productCategoryID)).Return(true, nil)
	repoMock.ProductCategoryRepository.Mock.On("Delete", int64(productCategoryID)).Return(nil)

	productCategory, appErr := service.Delete(int64(productCategoryID))

	assert.NotNil(t, productCategory)
	assert.Nil(t, appErr)
}

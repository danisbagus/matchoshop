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

var mockProductCategoryRepo = &mocks.ProductCategoryRepo{Mock: mock.Mock{}}
var productCategoryService = ProductCategoryService{repo: mockProductCategoryRepo}

func TestProductCategory_Create_NotValidated(t *testing.T) {

	req := dto.CreateProductCategoryRequest{}

	productCategory, appErr := productCategoryService.Create(&req)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Create_NameExits(t *testing.T) {

	req := dto.CreateProductCategoryRequest{
		Name: "Electronic",
	}

	mockProductCategoryRepo.Mock.On("CheckByName", req.Name).Return(true, nil)

	productCategory, appErr := productCategoryService.Create(&req)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Create_Success(t *testing.T) {

	req := dto.CreateProductCategoryRequest{
		Name: "Sport",
	}

	mockProductCategoryRepo.Mock.On("CheckByName", req.Name).Return(false, nil)

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
	mockProductCategoryRepo.Mock.On("Insert", &formProductCategory).Return(&resultProductCategory, nil)

	productCategory, appErr := productCategoryService.Create(&req)

	assert.NotNil(t, productCategory)
	assert.Nil(t, appErr)
}
func TestProductCategory_GetDetail_NotFound(t *testing.T) {

	var productCategoryID int64 = 10

	mockProductCategoryRepo.Mock.On("GetOneByID", productCategoryID).Return(nil, errs.NewNotFoundError("Product category not found!"))

	productCategory, appErr := productCategoryService.GetDetail(productCategoryID)
	assert.NotNil(t, appErr)
	assert.Nil(t, productCategory)
}

func TestProductCategory_GetDetail_Success(t *testing.T) {

	var productCategoryID int64 = 1

	productCategoryResult := domain.ProductCategory{
		ProductCategoryID: 1,
		Name:              "Modern shoes",
	}

	mockProductCategoryRepo.Mock.On("GetOneByID", productCategoryID).Return(&productCategoryResult, nil)

	productCategory, appErr := productCategoryService.GetDetail(productCategoryID)
	assert.NotNil(t, productCategory)
	assert.Nil(t, appErr)
}

func TestProductCategory_Update_NotValidated(t *testing.T) {

	productCategoryID := 1
	reqProductCategory := dto.CreateProductCategoryRequest{}

	productCategory, appErr := productCategoryService.Update(int64(productCategoryID), &reqProductCategory)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Update_NotFound(t *testing.T) {

	productCategoryID := 1
	req := dto.CreateProductCategoryRequest{
		Name: "Electonics",
	}

	mockProductCategoryRepo.Mock.On("CheckByID", int64(productCategoryID)).Return(false, nil)

	productCategory, appErr := productCategoryService.Update(int64(productCategoryID), &req)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Update_NameExits(t *testing.T) {

	productCategoryID := 2
	req := dto.CreateProductCategoryRequest{
		Name: "Electonics",
	}

	mockProductCategoryRepo.Mock.On("CheckByID", int64(productCategoryID)).Return(true, nil)

	mockProductCategoryRepo.Mock.On("CheckByIDAndName", int64(productCategoryID), req.Name).Return(true, nil)

	productCategory, appErr := productCategoryService.Update(int64(productCategoryID), &req)

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Updated_Success(t *testing.T) {

	productCategoryID := 2
	req := dto.CreateProductCategoryRequest{
		Name: "Sport",
	}

	mockProductCategoryRepo.Mock.On("CheckByID", int64(productCategoryID)).Return(true, nil)

	mockProductCategoryRepo.Mock.On("CheckByIDAndName", int64(productCategoryID), req.Name).Return(false, nil)

	formProductCategory := domain.ProductCategory{
		Name:      req.Name,
		CreatedAt: time.Now().Format(dbTSLayout),
		UpdatedAt: time.Now().Format(dbTSLayout),
	}

	mockProductCategoryRepo.Mock.On("Update", int64(productCategoryID), &formProductCategory).Return(nil)

	productCategory, appErr := productCategoryService.Update(int64(productCategoryID), &req)

	assert.NotNil(t, productCategory)
	assert.Nil(t, appErr)
}

func TestProductCategory_Delete_NotFound(t *testing.T) {

	productCategoryID := 1

	mockProductCategoryRepo.Mock.On("CheckByID", int64(productCategoryID)).Return(false, nil)

	productCategory, appErr := productCategoryService.Delete(int64(productCategoryID))

	assert.Nil(t, productCategory)
	assert.NotNil(t, appErr)
}

func TestProductCategory_Delete_Success(t *testing.T) {

	productCategoryID := 2

	mockProductCategoryRepo.Mock.On("CheckByID", int64(productCategoryID)).Return(true, nil)

	mockProductCategoryRepo.Mock.On("Delete", int64(productCategoryID)).Return(nil)

	productCategory, appErr := productCategoryService.Delete(int64(productCategoryID))

	assert.NotNil(t, productCategory)
	assert.Nil(t, appErr)
}

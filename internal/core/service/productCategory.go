package service

import (
	"fmt"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type ProductCategoryService struct {
	repo port.ProductCategoryRepo
}

func NewProductCategoryService(repo port.ProductCategoryRepo) port.ProductCategoryService {
	return &ProductCategoryService{
		repo: repo,
	}
}

func (r ProductCategoryService) Create(req *dto.CreateProductCategoryRequest) (*dto.ResponseData, *errs.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	checkProductCategory, appErr := r.repo.CheckByName(req.Name)
	if appErr != nil {
		return nil, appErr
	}

	if checkProductCategory {
		errorMessage := fmt.Sprintf("Product category with name %s is already exits", req.Name)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	formProductCategory := domain.ProductCategory{
		Name:      req.Name,
		CreatedAt: time.Now().Format(dbTSLayout),
		UpdatedAt: time.Now().Format(dbTSLayout),
	}

	newProductCategoryData, err := r.repo.Insert(&formProductCategory)
	if err != nil {
		return nil, err
	}

	response := dto.NewCreateProductCategoryResponse("Sucessfully create data", newProductCategoryData)

	return response, nil
}

func (r ProductCategoryService) GetList() (*dto.ResponseData, *errs.AppError) {

	productCategories, appErr := r.repo.GetAll()
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewGetProductCategoryListResponse("Successfully get data", productCategories)

	return response, nil
}

func (r ProductCategoryService) GetDetail(productCategoryID int64) (*dto.ResponseData, *errs.AppError) {

	productCategory, appErr := r.repo.GetOneByID(productCategoryID)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewGetProductCategoryDetailResponse("Successfully get data", productCategory)

	return response, nil
}

func (r ProductCategoryService) Update(productCategoryID int64, req *dto.CreateProductCategoryRequest) (*dto.ResponseData, *errs.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	checkProductCategory, appErr := r.repo.CheckByID(productCategoryID)
	if appErr != nil {
		return nil, appErr
	}

	if !checkProductCategory {
		return nil, errs.NewBadRequestError("Product category not found")
	}

	checkProductCategory, appErr = r.repo.CheckByIDAndName(productCategoryID, req.Name)
	if appErr != nil {
		return nil, appErr
	}

	if checkProductCategory {
		errorMessage := fmt.Sprintf("Product Category with name %s is already exits", req.Name)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	formProductCategory := domain.ProductCategory{
		Name:      req.Name,
		CreatedAt: time.Now().Format(dbTSLayout),
		UpdatedAt: time.Now().Format(dbTSLayout),
	}

	appErr = r.repo.Update(productCategoryID, &formProductCategory)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.GenerateResponseData("Successfully update data", map[string]string{})

	return response, nil
}

func (r ProductCategoryService) Delete(productCategoryID int64) (*dto.ResponseData, *errs.AppError) {

	checkProductCategory, appErr := r.repo.CheckByID(productCategoryID)
	if appErr != nil {
		return nil, appErr
	}

	if !checkProductCategory {
		return nil, errs.NewBadRequestError("Product category not found")
	}

	appErr = r.repo.Delete(productCategoryID)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.GenerateResponseData("Successfully delete data", map[string]string{})

	return response, nil
}

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
	repo port.IProductCategoryRepo
}

func NewProductCategoryService(repo port.IProductCategoryRepo) port.IProductCategoryService {
	return &ProductCategoryService{
		repo: repo,
	}
}

func (r ProductCategoryService) Create(req *dto.CreateProductCategoryRequest) (*dto.CreateProductCategoryResponse, *errs.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	checkProductCategory, appErr := r.repo.CheckByMerchantIDAndName(req.MerchantID, req.Name)
	if appErr != nil {
		return nil, appErr
	}

	if checkProductCategory {
		errorMessage := fmt.Sprintf("Product category with name %s is already exits", req.Name)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	formProductCategory := domain.ProductCategory{
		MerchantID: req.MerchantID,
		Name:       req.Name,
		CreatedAt:  time.Now().Format(dbTSLayout),
		UpdatedAt:  time.Now().Format(dbTSLayout),
	}

	newProductCategoryData, err := r.repo.Insert(&formProductCategory)
	if err != nil {
		return nil, err
	}

	response := dto.NewCreateProductCategoryResponse(newProductCategoryData)

	return response, nil
}

func (r ProductCategoryService) GetList(merchantID int64) (*dto.ProductCategoryListResponse, *errs.AppError) {

	productCategories, appErr := r.repo.GetAllByMerchantID(merchantID)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewGetProductCategoryListResponse(productCategories)

	return response, nil
}

func (r ProductCategoryService) GetDetail(productCategoryID int64, merchantID int64) (*dto.ProductCategoryResponse, *errs.AppError) {

	productCategory, appErr := r.repo.GetOneByIDAndMerchantID(productCategoryID, merchantID)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewGetProductCategoryDetailResponse(productCategory)

	return response, nil
}

func (r ProductCategoryService) Update(productCategoryID int64, req *dto.CreateProductCategoryRequest) *errs.AppError {

	appErr := req.Validate()
	if appErr != nil {
		return appErr
	}

	checkProductCategory, appErr := r.repo.CheckByIDAndMerchantID(productCategoryID, req.MerchantID)
	if appErr != nil {
		return appErr
	}

	if !checkProductCategory {
		return errs.NewBadRequestError("Product category not found")
	}

	checkProductCategory, appErr = r.repo.CheckByIDAndMerchantIDAndName(productCategoryID, req.MerchantID, req.Name)
	if appErr != nil {
		return appErr
	}

	if checkProductCategory {
		errorMessage := fmt.Sprintf("Product Category with name %s is already exits", req.Name)
		return errs.NewBadRequestError(errorMessage)
	}

	formProductCategory := domain.ProductCategory{
		MerchantID: req.MerchantID,
		Name:       req.Name,
		CreatedAt:  time.Now().Format(dbTSLayout),
		UpdatedAt:  time.Now().Format(dbTSLayout),
	}

	appErr = r.repo.Update(productCategoryID, &formProductCategory)
	if appErr != nil {
		return appErr
	}

	return nil
}

func (r ProductCategoryService) Delete(productCategoryID int64, merchantID int64) *errs.AppError {

	checkProductCategory, appErr := r.repo.CheckByIDAndMerchantID(productCategoryID, merchantID)
	if appErr != nil {
		return appErr
	}

	if !checkProductCategory {
		return errs.NewBadRequestError("Product category not found")
	}

	appErr = r.repo.Delete(productCategoryID)
	if appErr != nil {
		return appErr
	}

	return nil
}

package usecase

import (
	"fmt"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/danisbagus/matchoshop/internal/repository"
)

type IProductCategoryUsecase interface {
	Create(data *domain.CreateProductCategoryRequest) (*domain.ResponseData, *errs.AppError)
	GetList() ([]domain.ProductCategoryModel, *errs.AppError)
	GetDetail(productCategoryID int64) (*domain.ResponseData, *errs.AppError)
	Update(productCategoryID int64, data *domain.CreateProductCategoryRequest) (*domain.ResponseData, *errs.AppError)
	Delete(productCategoryID int64) (*domain.ResponseData, *errs.AppError)
}

type ProductCategoryUsecase struct {
	productCategoryRepo repository.IProductCategoryRepository
}

func NewProductCategoryUsecase(repository repository.RepositoryCollection) IProductCategoryUsecase {
	return &ProductCategoryUsecase{
		productCategoryRepo: repository.ProductCategoryRepository,
	}
}

func (r ProductCategoryUsecase) Create(req *domain.CreateProductCategoryRequest) (*domain.ResponseData, *errs.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	checkProductCategory, appErr := r.productCategoryRepo.CheckByName(req.Name)
	if appErr != nil {
		return nil, appErr
	}

	if checkProductCategory {
		errorMessage := fmt.Sprintf("Product category with name %s is already exits", req.Name)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	formProductCategory := domain.ProductCategoryModel{
		Name:      req.Name,
		CreatedAt: time.Now().Format(dbTSLayout),
		UpdatedAt: time.Now().Format(dbTSLayout),
	}

	newProductCategoryData, err := r.productCategoryRepo.Insert(&formProductCategory)
	if err != nil {
		return nil, err
	}

	response := domain.NewCreateProductCategoryResponse("Sucessfully create data", newProductCategoryData)

	return response, nil
}

func (r ProductCategoryUsecase) GetList() ([]domain.ProductCategoryModel, *errs.AppError) {
	productCategories, appErr := r.productCategoryRepo.GetAll()
	if appErr != nil {
		return nil, appErr
	}
	return productCategories, nil
}

func (r ProductCategoryUsecase) GetDetail(productCategoryID int64) (*domain.ResponseData, *errs.AppError) {

	productCategory, appErr := r.productCategoryRepo.GetOneByID(productCategoryID)
	if appErr != nil {
		return nil, appErr
	}

	response := domain.NewGetProductCategoryDetailResponse("Successfully get data", productCategory)

	return response, nil
}

func (r ProductCategoryUsecase) Update(productCategoryID int64, req *domain.CreateProductCategoryRequest) (*domain.ResponseData, *errs.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	checkProductCategory, appErr := r.productCategoryRepo.CheckByID(productCategoryID)
	if appErr != nil {
		return nil, appErr
	}

	if !checkProductCategory {
		return nil, errs.NewBadRequestError("Product category not found")
	}

	checkProductCategory, appErr = r.productCategoryRepo.CheckByIDAndName(productCategoryID, req.Name)
	if appErr != nil {
		return nil, appErr
	}

	if checkProductCategory {
		errorMessage := fmt.Sprintf("Product Category with name %s is already exits", req.Name)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	formProductCategory := domain.ProductCategoryModel{
		Name:      req.Name,
		CreatedAt: time.Now().Format(dbTSLayout),
		UpdatedAt: time.Now().Format(dbTSLayout),
	}

	appErr = r.productCategoryRepo.Update(productCategoryID, &formProductCategory)
	if appErr != nil {
		return nil, appErr
	}

	response := domain.GenerateResponseData("Successfully update data", map[string]string{})

	return response, nil
}

func (r ProductCategoryUsecase) Delete(productCategoryID int64) (*domain.ResponseData, *errs.AppError) {

	checkProductCategory, appErr := r.productCategoryRepo.CheckByID(productCategoryID)
	if appErr != nil {
		return nil, appErr
	}

	if !checkProductCategory {
		return nil, errs.NewBadRequestError("Product category not found")
	}

	appErr = r.productCategoryRepo.Delete(productCategoryID)
	if appErr != nil {
		return nil, appErr
	}

	response := domain.GenerateResponseData("Successfully delete data", map[string]string{})

	return response, nil
}

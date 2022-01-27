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

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	checkProductCategory, err := r.repo.CheckByMerchantIDAndName(req.MerchantID, req.Name)
	if err != nil {
		return nil, err
	}

	if checkProductCategory {
		errorMessage := fmt.Sprintf("Product Category with name %s is alrady exits", req.Name)
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

	productCategories, err := r.repo.GetAllByMerchantID(merchantID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetProductCategoryListResponse(productCategories)

	return response, nil
}

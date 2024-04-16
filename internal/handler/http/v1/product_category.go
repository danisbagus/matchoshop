package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/cmd/middleware"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/danisbagus/matchoshop/internal/domain/common/constants"
	"github.com/danisbagus/matchoshop/internal/usecase"
	"github.com/gorilla/mux"
)

type ProductCategoryHandler struct {
	productCategoryUsecase usecase.IProductCategoryUsecase
}

func NewProductCategoryHandler(r *mux.Router, usecaseCollection usecase.UsecaseCollection, APIMiddleware middleware.IAPIMiddleware) {
	handler := ProductCategoryHandler{
		productCategoryUsecase: usecaseCollection.ProductCategoryUsecase,
	}

	route := r.PathPrefix("/api/v1/product-category").Subrouter()
	route.HandleFunc("", handler.GetProductCategoryList).Methods(http.MethodGet)
	route.HandleFunc("/{product_category_id}", handler.GetProductCategoryDetail).Methods(http.MethodGet)

	adminRoute := r.PathPrefix("/api/v1/admin/product-category").Subrouter()
	adminRoute.Use(APIMiddleware.Authorization(), APIMiddleware.ACL(constants.AdminPermission))
	adminRoute.HandleFunc("", handler.CreateProductCategory).Methods(http.MethodPost)
	adminRoute.HandleFunc("/{product_category_id}", handler.UpdateProductCategory).Methods(http.MethodPut)
	adminRoute.HandleFunc("/{product_category_id}", handler.Delete).Methods(http.MethodDelete)
}

func (h ProductCategoryHandler) CreateProductCategory(w http.ResponseWriter, r *http.Request) {
	var request domain.CreateProductCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("Error while decoding create product category request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create product category")
		return
	}

	createData, appErr := h.productCategoryUsecase.Create(&request)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, createData)
}

func (h ProductCategoryHandler) GetProductCategoryList(w http.ResponseWriter, r *http.Request) {

	productCategories, appErr := h.productCategoryUsecase.GetList()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.NewGetProductCategoryListResponse("Successfully get data", productCategories)
	response.Write(w, http.StatusOK, res)
}

func (h ProductCategoryHandler) GetProductCategoryDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productCategoryID, _ := strconv.Atoi(vars["product_category_id"])

	productCategory, appErr := h.productCategoryUsecase.GetDetail(int64(productCategoryID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, productCategory)
}

func (h ProductCategoryHandler) UpdateProductCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productCategoryID, _ := strconv.Atoi(vars["product_category_id"])

	var request domain.CreateProductCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("Error while decoding update product category request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create product category")
		return
	}

	updateData, appErr := h.productCategoryUsecase.Update(int64(productCategoryID), &request)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, updateData)
}

func (h ProductCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productCategoryID, _ := strconv.Atoi(vars["product_category_id"])

	deleteData, appErr := h.productCategoryUsecase.Delete(int64(productCategoryID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, deleteData)
}

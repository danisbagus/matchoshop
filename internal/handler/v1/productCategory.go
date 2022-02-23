package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/gorilla/mux"
)

type ProductCategoryHandler struct {
	Service port.ProductCategoryService
}

func (rc ProductCategoryHandler) CrateProductCategory(w http.ResponseWriter, r *http.Request) {

	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	var request dto.CreateProductCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("Error while decoding create product category request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create product category")
		return
	}

	appErr = checkAuthorizeByRoleID(claimData.RoleID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	createData, appErr := rc.Service.Create(&request)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, createData)
}

func (rc ProductCategoryHandler) GetProductCategoryList(w http.ResponseWriter, r *http.Request) {

	productCategories, appErr := rc.Service.GetList()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, productCategories)
}

func (rc ProductCategoryHandler) GetProductCategoryDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productCategoryID, _ := strconv.Atoi(vars["product_category_id"])

	productCategory, appErr := rc.Service.GetDetail(int64(productCategoryID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, productCategory)
}

func (rc ProductCategoryHandler) UpdateProductCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productCategoryID, _ := strconv.Atoi(vars["product_category_id"])

	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	var request dto.CreateProductCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("Error while decoding update product category request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create product category")
		return
	}

	appErr = checkAuthorizeByRoleID(claimData.RoleID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	updateData, appErr := rc.Service.Update(int64(productCategoryID), &request)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, updateData)
}

func (rc ProductCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productCategoryID, _ := strconv.Atoi(vars["product_category_id"])

	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	appErr = checkAuthorizeByRoleID(claimData.RoleID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	deleteData, appErr := rc.Service.Delete(int64(productCategoryID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, deleteData)
}

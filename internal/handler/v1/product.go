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

type ProductHandler struct {
	Service port.ProductService
}

func (rc ProductHandler) CrateProduct(w http.ResponseWriter, r *http.Request) {

	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	var request dto.CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("Error while decoding create product request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create product")
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

func (rc ProductHandler) GetProductList(w http.ResponseWriter, r *http.Request) {

	products, appErr := rc.Service.GetList()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, products)
}

func (rc ProductHandler) GetProductDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["product_id"])

	productCategory, appErr := rc.Service.GetDetail(int64(productID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, productCategory)
}

func (rc ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["product_id"])

	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	var request dto.CreateProductRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("Error while decoding update product request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create product category")
		return
	}

	appErr = checkAuthorizeByRoleID(claimData.RoleID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	updateData, appErr := rc.Service.Update(int64(productID), &request)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, updateData)
}

func (rc ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productID, _ := strconv.Atoi(vars["product_id"])

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

	deleteData, appErr := rc.Service.Delete(int64(productID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, deleteData)
}

package v1

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type ProductCategoryHandler struct {
	Service port.IProductCategoryService
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

	if claimData.MerchantID != request.MerchantID {
		response.Error(w, http.StatusBadRequest, "Not allowed use not owned merchant ID")
		return
	}

	data, appErr := rc.Service.Create(&request)

	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, data)
}

func (rc ProductCategoryHandler) GetProductCategoryList(w http.ResponseWriter, r *http.Request) {

	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	productCategories, appErr := rc.Service.GetList(claimData.MerchantID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, productCategories)
}

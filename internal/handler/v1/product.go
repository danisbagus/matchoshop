package v1

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type ProductHandler struct {
	Service port.IProductService
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

	request.MerchantID = claimData.MerchantID

	createData, appErr := rc.Service.Create(&request)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, createData)
}

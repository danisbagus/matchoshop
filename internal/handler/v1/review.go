package v1

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/utils/auth"
	"github.com/danisbagus/matchoshop/utils/constants"
	"github.com/danisbagus/matchoshop/utils/helper"
	"github.com/gorilla/mux"
)

type ReviewHandler struct {
	Service port.ReviewService
}

func (h ReviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	var req dto.ReviewRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := req.Validate()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	form := new(domain.Review)
	form.UserID = userInfo.UserID
	form.ProductID = req.ProductID
	form.Rating = req.Rating
	form.Comment = req.Comment

	appErr = h.Service.Create(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	resData := dto.GenerateResponseData(constants.SuccessCreate, nil)
	response.Write(w, http.StatusCreated, resData)
}

func (h ReviewHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	vars := mux.Vars(r)
	productID := helper.StringToInt64(vars["product_id"], 0)

	review, appErr := h.Service.GetDetail(userInfo.UserID, productID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	resData := dto.NewReviewResponse(constants.SuccesGet, review)
	response.Write(w, http.StatusOK, resData)
}

func (h ReviewHandler) Update(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	var req dto.ReviewRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	appErr := req.Validate()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	form := new(domain.Review)
	form.UserID = userInfo.UserID
	form.ProductID = req.ProductID
	form.Rating = req.Rating
	form.Comment = req.Comment

	appErr = h.Service.Create(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	resData := dto.GenerateResponseData(constants.SuccessUpdate, nil)
	response.Write(w, http.StatusCreated, resData)
}

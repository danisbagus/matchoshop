package v1

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/utils/auth"
)

type UserHandler struct {
	Service port.UserService
}

func (rc UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed to login")
		return
	}

	token, appErr := rc.Service.Login(loginRequest)
	if appErr != nil {
		response.Write(w, appErr.Code, appErr.AsMessage())
		return
	}

	response.Write(w, http.StatusOK, *token)
}

func (rc UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var refreshRequest dto.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&refreshRequest); err != nil {
		logger.Error("Error while decoding refresh token request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed to refresh token")
		return
	}

	token, appErr := rc.Service.Refresh(refreshRequest)

	if appErr != nil {

		response.Write(w, appErr.Code, appErr.AsMessage())
		return
	}

	response.Write(w, http.StatusOK, *token)
}

func (rc UserHandler) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	var registerRequest dto.RegisterCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		logger.Error("Error while decoding register customer request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, appErr := rc.Service.RegisterCustomer(&registerRequest)
	if appErr != nil {
		response.Write(w, appErr.Code, appErr.AsMessage())
		return
	}

	response.Write(w, http.StatusOK, *token)
}

func (rc UserHandler) GetUserDetail(w http.ResponseWriter, r *http.Request) {

	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)

	userData, appErr := rc.Service.GetDetail(userInfo.UserID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, userData)
}

func (rc UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)

	var request dto.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error("Error while decoding update user request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create user category")
		return
	}

	updateData, appErr := rc.Service.Update(userInfo.UserID, &request)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, updateData)
}

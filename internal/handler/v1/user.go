package v1

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type AuthHandler struct {
	Service port.IAuthService
}

func (rc AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		response.Write(w, http.StatusBadRequest, "Failed to login")
		return
	}

	token, appErr := rc.Service.Login(loginRequest)
	if appErr != nil {
		response.Write(w, appErr.Code, appErr.AsMessage())
		return
	}

	response.Write(w, http.StatusOK, *token)
}

func (rc AuthHandler) RegisterMerchant(w http.ResponseWriter, r *http.Request) {
	var registerRequest dto.RegisterMerchantRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		logger.Error("Error while decoding register merchant request: " + err.Error())
		response.Write(w, http.StatusBadRequest, "Failed to register merchant")
		return
	}

	token, appErr := rc.Service.RegisterMerchant(&registerRequest)
	if appErr != nil {
		response.Write(w, appErr.Code, appErr.AsMessage())
		return
	}
	response.Write(w, http.StatusOK, *token)
}

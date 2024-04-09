package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danisbagus/go-common-packages/http/response"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/danisbagus/matchoshop/utils/auth"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service port.UserService
}

func (rc UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed to login: "+err.Error())
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
	var refreshRequest domain.RefreshTokenRequest
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
	var registerRequest domain.RegisterCustomerRequest
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
	req := new(domain.UpdateUserRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logger.Error("Error while decoding update user request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create user category")
		return
	}

	appErr := req.Validate()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	form := new(domain.UserModel)
	form.UserID = userInfo.UserID
	form.Name = req.Name

	appErr = rc.Service.Update(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.NewGetUserDetailResponse("Sucessfully update data", form)
	response.Write(w, http.StatusOK, res)
}

func (rc UserHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	roleID := userInfo.RoleID

	users, appErr := rc.Service.GetList(roleID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	resData := domain.NewGetUserListResponse("Successfully get data", users)
	response.Write(w, http.StatusOK, resData)
}

func (rc UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["user_id"])
	roleID := userInfo.RoleID

	appErr := rc.Service.Delete(int64(userID), roleID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.GenerateResponseData("Sucessfully delete data", nil)
	response.Write(w, http.StatusOK, res)
}

func (rc UserHandler) UpdateUserAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["user_id"])
	req := new(domain.UpdateUserRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		logger.Error("Error while decoding update user request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed create user category")
		return
	}

	appErr := req.Validate()
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	form := new(domain.UserModel)
	form.UserID = int64(userID)
	form.Name = req.Name

	appErr = rc.Service.Update(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.NewGetUserDetailResponse("Sucessfully update data", form)
	response.Write(w, http.StatusOK, res)
}

func (rc UserHandler) GetUserDetailAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["user_id"])
	userData, appErr := rc.Service.GetDetail(int64(userID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	response.Write(w, http.StatusOK, userData)
}

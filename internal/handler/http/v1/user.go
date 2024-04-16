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
	"github.com/danisbagus/matchoshop/utils/auth"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(r *mux.Router, usecaseCollection usecase.UsecaseCollection, APIMiddleware middleware.IAPIMiddleware) {
	handler := UserHandler{
		userUsecase: usecaseCollection.UserUsecase,
	}

	route := r.PathPrefix("/api/v1/user").Subrouter()
	route.Use(APIMiddleware.Authorization())
	route.HandleFunc("", handler.GetUserDetail).Methods(http.MethodGet)
	route.HandleFunc("/profile", handler.UpdateUser).Methods(http.MethodPatch)

	adminRoute := r.PathPrefix("/api/v1/admin/user").Subrouter()
	adminRoute.Use(APIMiddleware.Authorization(), APIMiddleware.ACL(constants.AdminPermission))
	adminRoute.HandleFunc("", handler.GetUserList).Methods(http.MethodGet)
	adminRoute.HandleFunc("/{user_id}", handler.GetUserDetailAdmin).Methods(http.MethodGet)
	adminRoute.HandleFunc("/{user_id}", handler.DeleteUser).Methods(http.MethodDelete)
	adminRoute.HandleFunc("/{user_id}", handler.UpdateUserAdmin).Methods(http.MethodPatch)

	authRoute := r.PathPrefix("/api/v1/auth").Subrouter()
	authRoute.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
	authRoute.HandleFunc("/refresh", handler.Refresh).Methods(http.MethodPost)
	authRoute.HandleFunc("/register/customer", handler.RegisterCustomer).Methods(http.MethodPost)
}

func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed to login: "+err.Error())
		return
	}

	token, appErr := h.userUsecase.Login(loginRequest)
	if appErr != nil {
		response.Write(w, appErr.Code, appErr.AsMessage())
		return
	}

	response.Write(w, http.StatusOK, *token)
}

func (h UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var refreshRequest domain.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&refreshRequest); err != nil {
		logger.Error("Error while decoding refresh token request: " + err.Error())
		response.Error(w, http.StatusBadRequest, "Failed to refresh token")
		return
	}

	token, appErr := h.userUsecase.Refresh(refreshRequest)

	if appErr != nil {

		response.Write(w, appErr.Code, appErr.AsMessage())
		return
	}

	response.Write(w, http.StatusOK, *token)
}

func (h UserHandler) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	var registerRequest domain.RegisterCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		logger.Error("Error while decoding register customer request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, appErr := h.userUsecase.RegisterCustomer(&registerRequest)
	if appErr != nil {
		response.Write(w, appErr.Code, appErr.AsMessage())
		return
	}

	response.Write(w, http.StatusOK, *token)
}

func (h UserHandler) GetUserDetail(w http.ResponseWriter, r *http.Request) {

	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)

	userData, appErr := h.userUsecase.GetDetail(userInfo.UserID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	response.Write(w, http.StatusOK, userData)
}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	appErr = h.userUsecase.Update(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.NewGetUserDetailResponse("Sucessfully update data", form)
	response.Write(w, http.StatusOK, res)
}

func (h UserHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	roleID := userInfo.RoleID

	users, appErr := h.userUsecase.GetList(roleID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	resData := domain.NewGetUserListResponse("Successfully get data", users)
	response.Write(w, http.StatusOK, resData)
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(*auth.AccessTokenClaims)
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["user_id"])
	roleID := userInfo.RoleID

	appErr := h.userUsecase.Delete(int64(userID), roleID)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.GenerateResponseData("Sucessfully delete data", nil)
	response.Write(w, http.StatusOK, res)
}

func (h UserHandler) UpdateUserAdmin(w http.ResponseWriter, r *http.Request) {
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

	appErr = h.userUsecase.Update(form)
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}

	res := domain.NewGetUserDetailResponse("Sucessfully update data", form)
	response.Write(w, http.StatusOK, res)
}

func (h UserHandler) GetUserDetailAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["user_id"])
	userData, appErr := h.userUsecase.GetDetail(int64(userID))
	if appErr != nil {
		response.Error(w, appErr.Code, appErr.Message)
		return
	}
	response.Write(w, http.StatusOK, userData)
}

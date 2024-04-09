package domain

import (
	"github.com/danisbagus/go-common-packages/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	UserModel struct {
		UserID    int64  `db:"user_id"`
		Email     string `db:"email"`
		Password  string `db:"password"`
		Name      string `db:"name"`
		RoleID    int64  `db:"role_id"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}

	UserDetail struct {
		UserModel
		RoleName string `db:"role_name"`
	}

	ResponseData struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RefreshTokenRequest struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	RegisterCustomerRequest struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	UpdateUserRequest struct {
		Name string `json:"name"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		UserID       int64  `json:"user_id"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		RoleID       int64  `json:"role_id"`
	}

	UserDetailResponse struct {
		UserID int64  `json:"user_id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		RoleID int64  `json:"role_id"`
	}

	UserListResponse struct {
		UserDetailResponse
		RoleName string `json:"role_name"`
	}
	RefreshTokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	RegisterCustomerResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		UserID       int64  `json:"user_id"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		RoleID       int64  `json:"role_id"`
	}
)

func GenerateResponseData(message string, data interface{}) *ResponseData {
	return &ResponseData{
		Message: message,
		Data:    data,
	}
}

func NewLoginResponse(message string, accessToken string, refreshToken string, user *UserModel) *ResponseData {

	loginResponse := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.UserID,
		Name:         user.Name,
		Email:        user.Email,
		RoleID:       user.RoleID,
	}

	return GenerateResponseData(message, loginResponse)
}

func NewRefreshTokenResponse(message string, accessToken string) *ResponseData {

	refreshTokenResponse := RefreshTokenResponse{
		AccessToken: accessToken,
	}

	return GenerateResponseData(message, refreshTokenResponse)
}

func NewRegisterUserCustomerResponse(message string, accessToken string, refreshToken string, user *UserModel) *ResponseData {

	registerResponse := RegisterCustomerResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.UserID,
		Name:         user.Name,
		Email:        user.Email,
		RoleID:       user.RoleID,
	}

	return GenerateResponseData(message, registerResponse)
}

func NewGetUserDetailResponse(message string, data *UserModel) *ResponseData {

	userDetailResponse := &UserDetailResponse{
		UserID: data.UserID,
		Name:   data.Name,
		Email:  data.Email,
		RoleID: data.RoleID,
	}

	return GenerateResponseData(message, userDetailResponse)
}

func NewGetUserListResponse(message string, data []UserDetail) *ResponseData {
	users := make([]UserListResponse, 0)

	for _, value := range data {
		var user UserListResponse
		user.UserID = value.UserID
		user.Name = value.Name
		user.Email = value.Email
		user.RoleID = value.RoleID
		user.RoleName = value.RoleName

		users = append(users, user)
	}

	return GenerateResponseData(message, users)
}

func (r LoginRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Email, validation.Required); err != nil {
		return errs.NewBadRequestError("Email is required")
	} else if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errs.NewBadRequestError("Password is required")
	}

	return nil
}

func (r RegisterCustomerRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Name, validation.Required); err != nil {
		return errs.NewBadRequestError("Name is required")
	} else if err := validation.Validate(r.Email, validation.Required); err != nil {
		return errs.NewBadRequestError("Email is required")
	} else if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errs.NewBadRequestError("Password is required")
	} else if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errs.NewBadRequestError("Password is required")
	} else if r.Password != r.ConfirmPassword {
		return errs.NewBadRequestError("Invalid confirm password")
	}

	return nil
}

func (r UpdateUserRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Name, validation.Required); err != nil {
		return errs.NewBadRequestError("name is required")
	}
	return nil
}

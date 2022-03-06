package dto

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterCustomerRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       int64  `json:"user_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	RoleID       int64  `json:"role_id"`
}

type UserDetailResponse struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID int64  `json:"role_id"`
}
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterCustomerResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       int64  `json:"user_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	RoleID       int64  `json:"role_id"`
}

func GenerateResponseData(message string, data interface{}) *ResponseData {
	return &ResponseData{
		Message: message,
		Data:    data,
	}
}

func NewLoginResponse(message string, accessToken string, refreshToken string, user *domain.User) *ResponseData {

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

func NewRegisterUserCustomerResponse(message string, accessToken string, refreshToken string, user *domain.User) *ResponseData {

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

func NewGetUserDetailResponse(message string, data *domain.User) *ResponseData {

	userDetailResponse := &UserDetailResponse{
		UserID: data.UserID,
		Name:   data.Name,
		Email:  data.Email,
		RoleID: data.RoleID,
	}

	return GenerateResponseData(message, userDetailResponse)
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

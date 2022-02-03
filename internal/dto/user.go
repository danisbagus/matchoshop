package dto

import (
	"github.com/danisbagus/go-common-packages/errs"
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

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateResponseData(message string, data interface{}) *ResponseData {
	return &ResponseData{
		Message: message,
		Data:    data,
	}
}

func NewLoginResponse(message string, accessToken string, refreshToken string) *ResponseData {

	loginResponse := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return GenerateResponseData(message, loginResponse)
}

func (r LoginRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Email, validation.Required); err != nil {
		return errs.NewBadRequestError("Email is required")
	} else if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errs.NewBadRequestError("Password is required")
	}

	return nil
}

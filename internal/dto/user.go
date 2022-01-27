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

type RegisterMerchantRequest struct {
	Email              string `json:"email"`
	Password           string `json:"password"`
	MerchantName       string `json:"merchant_name"`
	MerchantIdentifier string `json:"merchant_identifier"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type RegisterMerchantResponse struct {
	UserID             int64  `json:"user_id"`
	MerchantID         int64  `json:"merchant_id"`
	MerchantName       string `json:"merchant_name"`
	MerchantIdentifier string `json:"merchant_identifier"`
}

func GenerateResponseData(message string, data interface{}) *ResponseData {
	return &ResponseData{
		Message: message,
		Data:    data,
	}
}

func NewLoginResponse(message string, accessToken string) *ResponseData {

	loginResponse := LoginResponse{
		AccessToken: accessToken,
	}

	return GenerateResponseData(message, loginResponse)
}

func NewRegisterUserMerchantResponse(message string, data *domain.UserMerchant) *ResponseData {

	registerResponse := RegisterMerchantResponse{
		UserID:             data.UserID,
		MerchantID:         data.MerchantID,
		MerchantName:       data.MerchantName,
		MerchantIdentifier: data.MerchantIdentifier,
	}

	return GenerateResponseData(message, registerResponse)
}

func (r LoginRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Email, validation.Required); err != nil {
		return errs.NewBadRequestError("Email is required")
	} else if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errs.NewBadRequestError("Password is required")
	}

	return nil
}

func (r RegisterMerchantRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Email, validation.Required); err != nil {
		return errs.NewBadRequestError("email is required")
	} else if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errs.NewBadRequestError("Password is required")
	} else if err := validation.Validate(r.MerchantName, validation.Required); err != nil {
		return errs.NewBadRequestError("Merchant name is required")
	} else if err := validation.Validate(r.MerchantIdentifier, validation.Required); err != nil {
		return errs.NewBadRequestError("Merchant identifier is required")
	}

	return nil
}

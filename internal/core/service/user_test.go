package service

import (
	"testing"

	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockUserRepo = &mocks.UserRepo{Mock: mock.Mock{}}
var mockRefreshTokenStoreRepo = &mocks.RefreshTokenStoreRepo{Mock: mock.Mock{}}

var userService = UserService{repo: mockUserRepo, refreshTokenStoreRepo: mockRefreshTokenStoreRepo}

func TestUser_Login_NotValidated(t *testing.T) {
	// Arrange
	req := dto.LoginRequest{}

	// Act
	login, appErr := userService.Login(req)

	// Assert
	assert.Nil(t, login)
	assert.NotNil(t, appErr)
}

func TestUser_Login_NotFound(t *testing.T) {

	req := dto.LoginRequest{
		Email:    "edagangan@live.com",
		Password: "test12",
	}

	resultFindOne := domain.User{
		UserID: 0,
	}
	mockUserRepo.Mock.On("FindOne", req.Email).Return(&resultFindOne, nil)

	login, appErr := userService.Login(req)

	assert.Nil(t, login)
	assert.NotNil(t, appErr)
}

func TestUser_Login_PasswordNotMatch(t *testing.T) {

	req := dto.LoginRequest{
		Email:    "matcho@live.com",
		Password: "test12",
	}

	resultFindOne := domain.User{
		Email:    "matcho@live.com",
		Password: "$2a$14$jCf0U5Ic9QI7RZdYGmgABOlV27nYac7Xg5iwoby/HFW.lcU8xqvaW",
	}

	mockUserRepo.Mock.On("FindOne", req.Email).Return(&resultFindOne, nil)

	login, appErr := userService.Login(req)

	assert.Nil(t, login)
	assert.NotNil(t, appErr)
}

func TestUser_Login_Success(t *testing.T) {

	req := dto.LoginRequest{
		Email:    "matcho@live.com",
		Password: "test12345",
	}

	resultFindOne := domain.User{
		Email:    "matcho@live.com",
		Password: "$2a$14$jCf0U5Ic9QI7RZdYGmgABOlV27nYac7Xg5iwoby/HFW.lcU8xqvaW",
	}

	mockUserRepo.Mock.On("FindOne", req.Email).Return(&resultFindOne, nil)

	resAccessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlX2lkIjoxLCJleHAiOjE2NDM5NzQ5MTh9.jfOtv_66VPWzGQRY1ZPSsMzPglUAjVLYkdMqRG1WQXA"
	resRefreshToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbl90eXBlIjoicmVmcmVzaF90b2tlbiIsInVzZXJfaWQiOjEsInJvbGVfaWQiOjEsImV4cCI6MTY0NjU2MzMxOH0.hTDIKmmlbEqqhNWrZua9osl-P6ftzzNe8MOcjk4ZUCg"

	mockUserRepo.Mock.On("GenerateAccessTokenAndRefreshToken", &resultFindOne).Return(resAccessToken, resRefreshToken, nil)

	mockRefreshTokenStoreRepo.Mock.On("Insert", resRefreshToken).Return(nil)

	login, appErr := userService.Login(req)

	assert.NotNil(t, login)
	assert.Nil(t, appErr)
}

func TestUser_Register_Not_Validated(t *testing.T) {

	req := dto.RegisterCustomerRequest{
		Email:           "matcho@live.com",
		Name:            "customer 1",
		Password:        "test12345",
		ConfirmPassword: "test123456",
	}

	register, appErr := userService.RegisterCustomer(&req)

	assert.Nil(t, register)
	assert.NotNil(t, appErr)
}

func TestUser_Register_Email_Already_Used(t *testing.T) {

	req := dto.RegisterCustomerRequest{
		Email:           "matcho@live.com",
		Name:            "customer 1",
		Password:        "test12345",
		ConfirmPassword: "test123456",
	}

	resultFindOne := domain.User{
		UserID: 1,
	}
	mockUserRepo.Mock.On("FindOne", req.Email).Return(&resultFindOne, nil)

	register, appErr := userService.RegisterCustomer(&req)

	assert.Nil(t, register)
	assert.NotNil(t, appErr)
}

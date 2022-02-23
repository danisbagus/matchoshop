package service

import (
	"testing"

	"github.com/danisbagus/go-common-packages/errs"
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

	mockUserRepo.Mock.On("FindOne", req.Email).Return(nil, errs.NewAuthenticationError("invalid credentials"))

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

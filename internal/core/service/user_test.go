package service

import (
	"testing"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockUserRepo = &mocks.IUserRepo{Mock: mock.Mock{}}
var userService = UserService{repo: mockUserRepo}

func TestUser_Login_FindOne_DBError(t *testing.T) {

	email := ""
	mockUserRepo.Mock.On("FindOne", email).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	login, appErr := userService.repo.FindOne(email)
	assert.NotNil(t, appErr)
	assert.Nil(t, login)
}

func TestUser_Login_FindOne_NotFound(t *testing.T) {

	email := "edagangan@live.com"
	mockUserRepo.Mock.On("FindOne", email).Return(nil, errs.NewAuthenticationError("invalid credentials"))

	login, appErr := userService.repo.FindOne(email)
	assert.NotNil(t, appErr)
	assert.Nil(t, login)
}

func TestUser_Login_FindOne_Success(t *testing.T) {

	email := "matcho@live.com"

	loginResult := domain.User{
		UserID:   1,
		Email:    "matcho@live.com",
		Password: "passwd",
		RoleID:   1,
	}

	mockUserRepo.Mock.On("FindOne", email).Return(&loginResult, nil)

	login, appErr := userService.repo.FindOne(email)
	assert.NotNil(t, login)
	assert.Nil(t, appErr)
}

func TestUser_Login_CheckPasswordHash_False(t *testing.T) {

	// Arrange
	reqPassword := "test12"
	loginPassword := "$2a$14$jCf0U5Ic9QI7RZdYGmgABOlV27nYac7Xg5iwoby/HFW.lcU8xqvaW"

	// Act
	match := checkPasswordHash(reqPassword, loginPassword)

	// Assert
	assert.Equal(t, match, false)
}

func TestUser_Login_CheckPasswordHash_True(t *testing.T) {

	reqPassword := "test12345"
	loginPassword := "$2a$14$jCf0U5Ic9QI7RZdYGmgABOlV27nYac7Xg5iwoby/HFW.lcU8xqvaW"

	match := checkPasswordHash(reqPassword, loginPassword)

	assert.Equal(t, match, true)
}

func TestUser_Login_CheckPasswordHash_NewAccessToken_Success(t *testing.T) {
	login := domain.User{
		UserID:   1,
		Password: "test12345",
		Email:    "edagangan@live.com",
		RoleID:   1,
	}

	claims := login.ClaimsForAccessToken()

	authToken := domain.NewAuthToken(claims)
	accessToken, appErr := authToken.NewAccessToken()

	assert.NotNil(t, accessToken)
	assert.Nil(t, appErr)
}

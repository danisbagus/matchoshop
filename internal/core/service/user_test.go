package service

import (
	"testing"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/internal/repository"
	"github.com/danisbagus/matchoshop/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func setupUserTest(t *testing.T) (mocks.RepoCollectionMocks, port.UserService) {
	repoMock := mocks.RepoCollectionMocks{
		UserRepository: mocks.NewIUserRepository(t),
	}

	repoCollection := repository.RepositoryCollection{
		UserRepository: repoMock.UserRepository,
	}

	service := NewUserService(repoCollection)
	return repoMock, service
}

func TestUser_Login_NotValidated(t *testing.T) {
	_, service := setupUserTest(t)

	// Arrange
	req := dto.LoginRequest{}

	// Act
	login, appErr := service.Login(req)

	// Assert
	assert.Nil(t, login)
	assert.NotNil(t, appErr)
}

func TestUser_Login_NotFound(t *testing.T) {
	repoMock, service := setupUserTest(t)

	req := dto.LoginRequest{
		Email:    "edagangan@live.com",
		Password: "test12",
	}

	resultFindOne := domain.User{
		UserID: 0,
	}

	repoMock.UserRepository.Mock.On("FindOne", req.Email).Return(&resultFindOne, nil)

	login, appErr := service.Login(req)

	assert.Nil(t, login)
	assert.NotNil(t, appErr)
}

func TestUser_Login_PasswordNotMatch(t *testing.T) {
	repoMock, service := setupUserTest(t)

	req := dto.LoginRequest{
		Email:    "matcho@live.com",
		Password: "test12",
	}

	resultFindOne := domain.User{
		Email:    "matcho@live.com",
		Password: "$2a$14$jCf0U5Ic9QI7RZdYGmgABOlV27nYac7Xg5iwoby/HFW.lcU8xqvaW",
	}

	repoMock.UserRepository.Mock.On("FindOne", req.Email).Return(&resultFindOne, nil)

	login, appErr := service.Login(req)

	assert.Nil(t, login)
	assert.NotNil(t, appErr)
}

func TestUser_Register_Not_Validated(t *testing.T) {
	_, service := setupUserTest(t)

	req := dto.RegisterCustomerRequest{
		Email:           "matcho@live.com",
		Name:            "customer 1",
		Password:        "test12345",
		ConfirmPassword: "test123456",
	}

	register, appErr := service.RegisterCustomer(&req)

	assert.Nil(t, register)
	assert.NotNil(t, appErr)
}

func TestUser_Register_Email_Already_Used(t *testing.T) {
	repoMock, service := setupUserTest(t)

	req := dto.RegisterCustomerRequest{
		Email:           "matcho@live.com",
		Name:            "customer 1",
		Password:        "test12345",
		ConfirmPassword: "test123456",
	}

	resultFindOne := domain.User{
		UserID: 1,
	}
	repoMock.UserRepository.Mock.On("FindOne", req.Email).Return(&resultFindOne, nil)

	register, appErr := service.RegisterCustomer(&req)

	assert.Nil(t, register)
	assert.NotNil(t, appErr)
}

func TestUser_Update_Not_Validated(t *testing.T) {
	_, service := setupUserTest(t)

	form := new(domain.User)
	form.UserID = 2

	appErr := service.Update(form)

	assert.NotNil(t, appErr)
}

func TestUser_Update_User_Not_Found(t *testing.T) {
	repoMock, service := setupUserTest(t)

	form := new(domain.User)
	form.Name = "Customer 3"
	form.UserID = 2

	resFindOneByID := domain.User{
		UserID: 0,
	}

	repoMock.UserRepository.Mock.On("FindOneById", form.UserID).Return(&resFindOneByID, nil)

	appErr := service.Update(form)

	assert.NotNil(t, appErr)
}

func TestUser_Update_Unexpected_Error_Update(t *testing.T) {
	repoMock, service := setupUserTest(t)

	form := new(domain.User)
	form.Name = "Customer 112345678901234567890123456789012345678901234567890234567890"
	form.UserID = 1

	resFindOneByID := domain.User{
		UserID: 1,
		Name:   "Customer 3",
	}

	repoMock.UserRepository.Mock.On("FindOneById", form.UserID).Return(&resFindOneByID, nil)

	formUpdate := domain.User{
		Name: form.Name,
	}

	repoMock.UserRepository.Mock.On("Update", form.UserID, &formUpdate).Return(errs.NewUnexpectedError("Unexpected database error"))
	appErr := service.Update(form)

	assert.NotNil(t, appErr)
}

func TestUser_Update_Success(t *testing.T) {
	repoMock, service := setupUserTest(t)

	form := new(domain.User)
	form.Name = "Customer 4"
	form.UserID = 1

	resFindOneByID := domain.User{
		UserID: 1,
		Name:   "Customer 4",
	}

	repoMock.UserRepository.Mock.On("FindOneById", form.UserID).Return(&resFindOneByID, nil)

	formUpdate := domain.User{
		Name: form.Name,
	}

	repoMock.UserRepository.Mock.On("Update", form.UserID, &formUpdate).Return(nil)
	appErr := service.Update(form)

	assert.Nil(t, appErr)
}

func TestUser_GetDetail_UserNotFound(t *testing.T) {
	repoMock, service := setupUserTest(t)

	userID := 2

	resFindOneByID := &domain.User{
		UserID: 0,
	}

	repoMock.UserRepository.Mock.On("FindOneById", int64(userID)).Return(resFindOneByID, nil)

	userDetail, appErr := service.GetDetail(int64(userID))

	assert.Nil(t, userDetail)
	assert.NotNil(t, appErr)
}

func TestUser_GetDetail_Success(t *testing.T) {
	repoMock, service := setupUserTest(t)

	userID := 2

	resFindOneByID := &domain.User{
		UserID: 2,
	}

	repoMock.UserRepository.Mock.On("FindOneById", int64(userID)).Return(resFindOneByID, nil)

	userDetail, appErr := service.GetDetail(int64(userID))

	assert.NotNil(t, userDetail)
	assert.Nil(t, appErr)
}

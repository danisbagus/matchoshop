package usecase

import (
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/domain"
	"github.com/danisbagus/matchoshop/internal/domain/common/constants"
	"github.com/danisbagus/matchoshop/internal/repository"
	"github.com/danisbagus/matchoshop/utils/auth"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const dbTSLayout = "2006-01-02 15:04:05"

type IUserUsecase interface {
	Login(req domain.LoginRequest) (*domain.ResponseData, *errs.AppError)
	Refresh(request domain.RefreshTokenRequest) (*domain.ResponseData, *errs.AppError)
	RegisterCustomer(req *domain.RegisterCustomerRequest) (*domain.ResponseData, *errs.AppError)
	GetList(roldID int64) ([]domain.UserDetail, *errs.AppError)
	GetDetail(userID int64) (*domain.ResponseData, *errs.AppError)
	Update(form *domain.UserModel) *errs.AppError
	Delete(userID, roleID int64) *errs.AppError
}

type UserUsecase struct {
	userRepo              repository.IUserRepository
	refreshTokenStoreRepo repository.IRefreshTokenStoreRepository
}

func NewUserUsecase(repository repository.RepositoryCollection) IUserUsecase {
	return &UserUsecase{
		userRepo:              repository.UserRepository,
		refreshTokenStoreRepo: repository.RefreshTokenStoreRepository,
	}
}

func (r UserUsecase) Login(req domain.LoginRequest) (*domain.ResponseData, *errs.AppError) {
	var appErr *errs.AppError
	var login *domain.UserModel

	appErr = req.Validate()

	if appErr != nil {
		return nil, appErr
	}

	login, appErr = r.userRepo.FindOne(req.Email)
	if appErr != nil {
		return nil, appErr
	}

	if login.UserID == 0 {
		return nil, errs.NewAuthenticationError("user not found")
	}

	match := checkPasswordHash(req.Password, login.Password)
	if !match {
		return nil, errs.NewAuthenticationError("invalid credentials")
	}

	accessToken, refreshToken, appErr := auth.GenerateAccessTokenAndRefreshToken(login.UserID, login.RoleID)
	if appErr != nil {
		logger.Error("Failed while generate access and refresh token: " + appErr.Message)
		return nil, appErr
	}

	appErr = r.refreshTokenStoreRepo.Insert(refreshToken)
	if appErr != nil {
		return nil, appErr
	}

	response := domain.NewLoginResponse("Successfully login", accessToken, refreshToken, login)

	return response, nil
}

func (r UserUsecase) Refresh(request domain.RefreshTokenRequest) (*domain.ResponseData, *errs.AppError) {

	// check token is valid or not
	validationErr := auth.IsTokenValid(request.AccessToken)
	if validationErr != nil {

		// check token is expired or not
		if validationErr.Errors == jwt.ValidationErrorExpired {
			var appErr *errs.AppError

			// check refresh token is exits or not
			checkRefreshToken, appErr := r.refreshTokenStoreRepo.CheckRefreshToken(request.RefreshToken)
			if appErr != nil {
				return nil, appErr
			}

			if !checkRefreshToken {
				return nil, errs.NewAuthenticationError("refresh token not found")
			}

			// generate a access token from refresh token.
			var accessToken string
			if accessToken, appErr = auth.NewAccessTokenFromRefreshToken(request.RefreshToken); appErr != nil {
				return nil, appErr
			}

			response := domain.NewRefreshTokenResponse("Successfully refresh token", accessToken)

			return response, nil
		}
		return nil, errs.NewAuthenticationError("invalid token")
	}

	logger.Error("Error while validate token")
	return nil, errs.NewAuthenticationError("cannot generate a new access token until the current one expires")
}

func (r UserUsecase) RegisterCustomer(req *domain.RegisterCustomerRequest) (*domain.ResponseData, *errs.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	// hashing password
	hashPassword, _ := hashPassword(req.Password)

	// validate email
	user, appErr := r.userRepo.FindOne(req.Email)
	if appErr != nil {
		return nil, appErr
	}

	if user.UserID != 0 {
		return nil, errs.NewAuthenticationError("Email already used")
	}

	form := domain.UserModel{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashPassword,
		RoleID:    3,
		CreatedAt: time.Now().Format(dbTSLayout),
		UpdatedAt: time.Now().Format(dbTSLayout),
	}

	newData, appErr := r.userRepo.CreateUserCustomer(&form)
	if appErr != nil {
		return nil, appErr
	}

	// generate access token and refresh token
	accessToken, refreshToken, appErr := auth.GenerateAccessTokenAndRefreshToken(newData.UserID, newData.RoleID)
	if appErr != nil {
		logger.Error("Failed while generate access and refresh token: " + appErr.Message)
		return nil, appErr

	}

	// insert refresh token
	appErr = r.refreshTokenStoreRepo.Insert(refreshToken)
	if appErr != nil {
		return nil, appErr
	}

	response := domain.NewRegisterUserCustomerResponse("Successfully register", accessToken, refreshToken, newData)

	return response, nil
}

func (r UserUsecase) GetDetail(userID int64) (*domain.ResponseData, *errs.AppError) {
	// get detail user
	userDetail, appErr := r.userRepo.FindOneById(userID)
	if appErr != nil {
		return nil, appErr
	}

	if userDetail.UserID == 0 {
		return nil, errs.NewAuthenticationError("user not found")
	}

	response := domain.NewGetUserDetailResponse("Successfully get data", userDetail)

	return response, nil
}

func (r UserUsecase) Update(form *domain.UserModel) *errs.AppError {

	user, appErr := r.userRepo.FindOneById(form.UserID)
	if appErr != nil {
		return appErr
	}
	if user.UserID == 0 {
		return errs.NewBadRequestError("User not found")
	}

	form.UpdatedAt = time.Now().Format(dbTSLayout)
	appErr = r.userRepo.Update(form.UserID, form)
	if appErr != nil {
		return appErr
	}

	form.RoleID = user.RoleID
	form.Email = user.Email
	return nil
}

func (r UserUsecase) GetList(roleID int64) ([]domain.UserDetail, *errs.AppError) {
	result := make([]domain.UserDetail, 0)
	userList, appErr := r.userRepo.GetAll()
	if appErr != nil {
		return nil, appErr
	}

	if roleID == constants.SuperAdminRoleID {
		result = userList
	} else {
		for _, value := range userList {
			if value.RoleID == constants.SuperAdminRoleID {
				continue
			}
			var user domain.UserDetail
			user.UserID = value.UserID
			user.Name = value.Name
			user.Email = value.Email
			user.RoleID = value.RoleID
			user.RoleName = value.RoleName
		}
	}

	return result, nil
}

func (r UserUsecase) Delete(userID, roleID int64) *errs.AppError {

	user, appErr := r.userRepo.FindOneById(userID)
	if appErr != nil {
		return appErr
	}
	if user.UserID == 0 {
		return errs.NewBadRequestError("user not found")
	}

	if roleID >= user.RoleID {
		return errs.NewBadRequestError("not allowed delete this user")
	}

	appErr = r.userRepo.Delete(userID)
	if appErr != nil {
		return appErr
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package service

import (
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const dbTSLayout = "2006-01-02 15:04:05"

type UserService struct {
	repo                  port.UserRepo
	refreshTokenStoreRepo port.RefreshTokenStoreRepo
}

func NewUserService(repo port.UserRepo, refreshTokenStoreRepo port.RefreshTokenStoreRepo) port.UserService {
	return &UserService{
		repo:                  repo,
		refreshTokenStoreRepo: refreshTokenStoreRepo,
	}
}

func (r UserService) Login(req dto.LoginRequest) (*dto.ResponseData, *errs.AppError) {
	var appErr *errs.AppError
	var login *domain.User

	appErr = req.Validate()

	if appErr != nil {
		return nil, appErr
	}

	login, appErr = r.repo.FindOne(req.Email)
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

	accessToken, refreshToken, appErr := r.repo.GenerateAccessTokenAndRefreshToken(login)
	if appErr != nil {
		return nil, appErr
	}

	appErr = r.refreshTokenStoreRepo.Insert(refreshToken)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewLoginResponse("Successfully login", accessToken, refreshToken, login)

	return response, nil
}

func (r UserService) Refresh(request dto.RefreshTokenRequest) (*dto.ResponseData, *errs.AppError) {

	// check token is valid or not
	if validationErr := request.IsAccessTokenValid(); validationErr != nil {

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
			if accessToken, appErr = domain.NewAccessTokenFromRefreshToken(request.RefreshToken); appErr != nil {
				return nil, appErr
			}

			response := dto.NewRefreshTokenResponse("Successfully refresh token", accessToken)

			return response, nil
		}
		return nil, errs.NewAuthenticationError("invalid token")
	}

	return nil, errs.NewAuthenticationError("cannot generate a new access token until the current one expires")
}

func (r UserService) RegisterCustomer(req *dto.RegisterCustomerRequest) (*dto.ResponseData, *errs.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	// hashing password
	hashPassword, _ := hashPassword(req.Password)

	// validate email
	user, appErr := r.repo.FindOne(req.Email)
	if appErr != nil {
		return nil, appErr
	}

	if user.UserID != 0 {
		return nil, errs.NewAuthenticationError("Email already used")
	}

	form := domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashPassword,
		RoleID:    3,
		CreatedAt: time.Now().Format(dbTSLayout),
		UpdatedAt: time.Now().Format(dbTSLayout),
	}

	newData, appErr := r.repo.CreateUserCustomer(&form)
	if appErr != nil {
		return nil, appErr
	}

	// generate access token and refresh token
	accessToken, refreshToken, appErr := r.repo.GenerateAccessTokenAndRefreshToken(newData)
	if appErr != nil {
		return nil, appErr
	}

	// insert refresh token
	appErr = r.refreshTokenStoreRepo.Insert(refreshToken)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewRegisterUserCustomerResponse("Successfully register", accessToken, refreshToken, newData)

	return response, nil
}

func (r UserService) GetDetail(userID int64) (*dto.ResponseData, *errs.AppError) {
	// get detail user
	userDetail, appErr := r.repo.FindOneById(userID)
	if appErr != nil {
		return nil, appErr
	}

	if userDetail.UserID == 0 {
		return nil, errs.NewAuthenticationError("user not found")
	}

	response := dto.NewGetUserDetailResponse("Successfully get data", userDetail)

	return response, nil
}

func (r UserService) Update(userID int64, req *dto.UpdateUserRequest) (*dto.ResponseData, *errs.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	user, appErr := r.repo.FindOneById(userID)
	if appErr != nil {
		return nil, appErr
	}

	if user.UserID == 0 {
		return nil, errs.NewBadRequestError("User not found")
	}

	formUpdate := domain.User{
		Name: req.Name,
	}

	appErr = r.repo.Update(userID, &formUpdate)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.GenerateResponseData("Successfully update data", map[string]string{})

	return response, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package service

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const dbTSLayout = "2006-01-02 15:04:05"

type UserService struct {
	repo                  port.IUserRepo
	refreshTokenStoreRepo port.IRefreshTokenStoreRepo
}

func NewUserService(repo port.IUserRepo, refreshTokenStoreRepo port.IRefreshTokenStoreRepo) port.IUserService {
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

	if login, appErr = r.repo.FindOne(req.Email); appErr != nil {
		return nil, appErr
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

	response := dto.NewLoginResponse("Successfully login", accessToken, refreshToken)

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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package service

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AuthService struct {
	repo port.IUserRepo
}

func NewUserService(repo port.IUserRepo) port.IUserService {
	return &AuthService{
		repo: repo,
	}
}

func (r AuthService) Login(req dto.LoginRequest) (*dto.ResponseData, *errs.AppError) {
	var appErr *errs.AppError
	var login *domain.User

	err := req.Validate()

	if err != nil {
		return nil, err
	}

	if login, appErr = r.repo.FindOne(req.Email); appErr != nil {
		return nil, appErr
	}

	match := checkPasswordHash(req.Password, login.Password)
	if !match {
		return nil, errs.NewAuthenticationError("invalid credentials")
	}

	claims := login.ClaimsForAccessToken()

	authToken := domain.NewAuthToken(claims)

	var accessToken string
	if accessToken, appErr = authToken.NewAccessToken(); appErr != nil {
		return nil, appErr
	}

	response := dto.NewLoginResponse("Successfully login", accessToken)

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

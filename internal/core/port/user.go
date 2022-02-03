package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type IUserRepo interface {
	FindOne(email string) (*domain.User, *errs.AppError)
	Verify(token string) *errs.AppError
	GenerateAndSaveRefreshTokenToStore(authToken *domain.AuthToken) (string, *errs.AppError)
}

type IUserService interface {
	Login(req dto.LoginRequest) (*dto.ResponseData, *errs.AppError)
}

package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type IUserRepo interface {
	FindOne(email string) (*domain.User, *errs.AppError)
	Verify(token string) *errs.AppError
	GenerateAccessTokenAndRefreshToken(data *domain.User) (string, string, *errs.AppError)
}

type IUserService interface {
	Login(req dto.LoginRequest) (*dto.ResponseData, *errs.AppError)
	Refresh(request dto.RefreshTokenRequest) (*dto.ResponseData, *errs.AppError)
}

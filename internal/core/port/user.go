package port

import (
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
	"github.com/danisbagus/matchoshop/pkg/errs"
)

type IAuthRepo interface {
	FindOne(email string) (*domain.User, *errs.AppError)
	Verify(token string) *errs.AppError
	CreateUserMerchant(data *domain.UserMerchant) (*domain.UserMerchant, *errs.AppError)
}

type IAuthService interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
	RegisterMerchant(req *dto.RegisterMerchantRequest) (*dto.RegisterMerchantResponse, *errs.AppError)
}

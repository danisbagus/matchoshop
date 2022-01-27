package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type IUserRepo interface {
	FindOne(email string) (*domain.User, *errs.AppError)
	Verify(token string) *errs.AppError
	CreateUserMerchant(data *domain.UserMerchant) (*domain.UserMerchant, *errs.AppError)
}

type IUserService interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
	RegisterMerchant(req *dto.RegisterMerchantRequest) (*dto.RegisterMerchantResponse, *errs.AppError)
}

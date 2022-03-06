package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type UserRepo interface {
	FindOne(email string) (*domain.User, *errs.AppError)
	FindOneById(userID int64) (*domain.User, *errs.AppError)
	CreateUserCustomer(data *domain.User) (*domain.User, *errs.AppError)
	Update(userID int64, data *domain.User) *errs.AppError
}

type UserService interface {
	Login(req dto.LoginRequest) (*dto.ResponseData, *errs.AppError)
	Refresh(request dto.RefreshTokenRequest) (*dto.ResponseData, *errs.AppError)
	RegisterCustomer(req *dto.RegisterCustomerRequest) (*dto.ResponseData, *errs.AppError)
	GetDetail(userID int64) (*dto.ResponseData, *errs.AppError)
	Update(userID int64, data *dto.UpdateUserRequest) (*dto.ResponseData, *errs.AppError)
}

package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type UserRepo interface {
	GetAll() ([]domain.UserDetail, *errs.AppError)
	FindOne(email string) (*domain.User, *errs.AppError)
	FindOneById(userID int64) (*domain.User, *errs.AppError)
	CreateUserCustomer(data *domain.User) (*domain.User, *errs.AppError)
	Update(userID int64, data *domain.User) *errs.AppError
	Delete(userID int64) *errs.AppError
}

type UserService interface {
	Login(req dto.LoginRequest) (*dto.ResponseData, *errs.AppError)
	Refresh(request dto.RefreshTokenRequest) (*dto.ResponseData, *errs.AppError)
	RegisterCustomer(req *dto.RegisterCustomerRequest) (*dto.ResponseData, *errs.AppError)
	GetList(roldID int64) ([]domain.UserDetail, *errs.AppError)
	GetDetail(userID int64) (*dto.ResponseData, *errs.AppError)
	Update(form *domain.User) *errs.AppError
	Delete(userID, roleID int64) *errs.AppError
}

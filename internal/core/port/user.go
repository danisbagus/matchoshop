package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/domain"
)

type UserService interface {
	Login(req domain.LoginRequest) (*domain.ResponseData, *errs.AppError)
	Refresh(request domain.RefreshTokenRequest) (*domain.ResponseData, *errs.AppError)
	RegisterCustomer(req *domain.RegisterCustomerRequest) (*domain.ResponseData, *errs.AppError)
	GetList(roldID int64) ([]domain.UserDetail, *errs.AppError)
	GetDetail(userID int64) (*domain.ResponseData, *errs.AppError)
	Update(form *domain.UserModel) *errs.AppError
	Delete(userID, roleID int64) *errs.AppError
}

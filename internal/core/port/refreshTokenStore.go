package port

import (
	"github.com/danisbagus/go-common-packages/errs"
)

type RefreshTokenStoreRepo interface {
	Insert(refreshToken string) *errs.AppError
	CheckRefreshToken(refreshToken string) (bool, *errs.AppError)
}

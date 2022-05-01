package port

import (
	"mime/multipart"

	"github.com/danisbagus/go-common-packages/errs"
)

type UploadService interface {
	UploadImage(file multipart.File) (string, *errs.AppError)
}

package usecase

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/utils/helper"
)

type IUploadUsecase interface {
	UploadImage(file multipart.File) (string, *errs.AppError)
}

type UploadUsecase struct {
}

func NewUploadUsecase() IUploadUsecase {
	return &UploadUsecase{}
}

func (s UploadUsecase) UploadImage(file multipart.File) (string, *errs.AppError) {
	//create cloudinary instance
	cld, err := cloudinary.NewFromURL(helper.EnvCloudURL())

	if err != nil {
		return "", errs.NewUnexpectedError(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: helper.EnvCloudUploadFolder(),
	})

	if err != nil {
		return "", errs.NewUnexpectedError(err.Error())
	}

	return result.URL, nil
}

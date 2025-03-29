package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/utils/helper"
)

type UploadService struct {
}

func NewUploadService() port.UploadService {
	return &UploadService{}
}

func (s UploadService) UploadImage(file *multipart.FileHeader) (string, *errs.AppError) {
	src, err := file.Open()
	if err != nil {
		return "", errs.NewUnexpectedError(err.Error())
	}

	defer src.Close()

	//create cloudinary instance
	cloudUrl := helper.EnvCloudURL()
	if cloudUrl == "" {
		return "", errs.NewUnexpectedError("cloudinary url not found")
	}

	cld, err := cloudinary.NewFromURL(cloudUrl)

	if err != nil {
		return "", errs.NewUnexpectedError(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{
		Folder: helper.EnvCloudUploadFolder(),
	})

	if err != nil {
		return "", errs.NewUnexpectedError(err.Error())
	}

	return result.URL, nil
}

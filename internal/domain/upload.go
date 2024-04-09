package domain

import (
	"mime/multipart"

	"github.com/danisbagus/go-common-packages/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UploadImageRequest struct {
	File multipart.File
}

type UploadImageRequestResponse struct {
	Url string `json:"url"`
}

func NewUploadImageResponse(message string, url string) *ResponseData {
	response := UploadImageRequestResponse{
		Url: url,
	}
	return GenerateResponseData(message, response)
}

func (r UploadImageRequest) Validate() *errs.AppError {
	if err := validation.Validate(r.File, validation.Required); err != nil {
		return errs.NewBadRequestError("file is required")
	}
	return nil
}

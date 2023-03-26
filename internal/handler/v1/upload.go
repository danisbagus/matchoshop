package v1

import (
	"net/http"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/labstack/echo/v4"
)

type UploadHandler struct {
	service port.UploadService
}

func NewUploadHandler(service port.UploadService) *UploadHandler {
	return &UploadHandler{service: service}
}

func (h UploadHandler) UploadImage(c echo.Context) error {
	//  Ensure our file does not exceed 5MB
	// r.Body = http.MaxBytesReader(w, r.Body, 5*1024*1024)

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error("Error while read upload image file: " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	url, appErr := h.service.UploadImage(file)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr.AsMessage())
	}

	resData := map[string]interface{}{
		"message": "Successfully upload image",
		"url":     url,
	}
	return c.JSON(http.StatusOK, resData)
}

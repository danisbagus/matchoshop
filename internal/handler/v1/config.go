package v1

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type ConfigHandler struct {
}

func (h ConfigHandler) GetConfig(c echo.Context) error {
	databaseURL := os.Getenv("DATABASE_URL")
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	cloudinaryUploadFolder := os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	paypalClientID := os.Getenv("PAYPAL_CLIENT_ID")

	resData := map[string]interface{}{
		"database_url":             databaseURL,
		"claudinary_url":           cloudinaryURL,
		"claudinary_upload_folder": cloudinaryUploadFolder,
		"paypal_client_id":         paypalClientID,
	}

	return c.JSON(http.StatusOK, resData)
}

func (h ConfigHandler) GetPaypalConfig(c echo.Context) error {
	paypalClientID := os.Getenv("PAYPAL_CLIENT_ID")
	resData := map[string]interface{}{
		"client_id": paypalClientID,
	}

	return c.JSON(http.StatusOK, resData)
}

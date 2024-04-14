package v1

import (
	"net/http"
	"os"

	"github.com/danisbagus/go-common-packages/http/response"
)

type ConfigHandler struct {
}

func (h ConfigHandler) GetConfig(w http.ResponseWriter, r *http.Request) {
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
	response.Write(w, http.StatusOK, resData)
}

func (h ConfigHandler) GetPaypalConfig(w http.ResponseWriter, r *http.Request) {
	paypalClientID := os.Getenv("PAYPAL_CLIENT_ID")
	resData := map[string]interface{}{
		"client_id": paypalClientID,
	}
	response.Write(w, http.StatusOK, resData)
}

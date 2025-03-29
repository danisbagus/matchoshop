package helper

import (
	"os"
)

func EnvCloudURL() string {
	return os.Getenv("CLOUDINARY_URL")
}

func EnvCloudUploadFolder() string {
	folder := os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	if folder == "" {
		folder = "matchoshop"
	}

	return folder
}

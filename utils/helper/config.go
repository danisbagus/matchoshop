package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvCloudURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_URL")
}

func EnvCloudUploadFolder() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}

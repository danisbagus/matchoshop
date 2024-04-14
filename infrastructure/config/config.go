package config

import (
	"time"

	"github.com/spf13/viper"
)

var APP_HOST string
var APP_PORT int
var APP_NAME string

var POSTGRES_HOST string
var POSTGRES_PORT int
var POSTGRES_USERNAME string
var POSTGRES_PASSWORD string
var POSTGRES_DATABASE string
var POSTGRES_SSL_MODE string

var DB_MAX_IDLE_CONNECTION int
var DB_MAX_OPEN_CONNECTION int
var DB_CONN_MAX_LIFETIME_SECONDS time.Duration

var CLOUDINARY_URL string
var CLOUDINARY_UPLOAD_FOLDER string

var PAYPAL_CLIENT_ID string

// this will reload config either from file or from system's ENV
// see: infrastructure/confi/setup.go
func reloadConfig() {
	APP_HOST = viper.GetString("APP_HOST")
	APP_PORT = viper.GetInt("APP_PORT")
	APP_NAME = viper.GetString("APP_NAME")

	POSTGRES_HOST = viper.GetString("POSTGRES_HOST")
	POSTGRES_PORT = viper.GetInt("POSTGRES_PORT")
	POSTGRES_USERNAME = viper.GetString("POSTGRES_USERNAME")
	POSTGRES_PASSWORD = viper.GetString("POSTGRES_PASSWORD")
	POSTGRES_DATABASE = viper.GetString("POSTGRES_DATABASE")
	POSTGRES_SSL_MODE = viper.GetString("POSTGRES_SSL_MODE")

	viper.SetDefault("DB_MAX_IDLE_CONNECTION", 10)
	viper.SetDefault("DB_MAX_OPEN_CONNECTION", 10)
	viper.SetDefault("DB_CONN_MAX_LIFETIME_SECONDS", time.Second*time.Duration(300))
	DB_MAX_IDLE_CONNECTION = viper.GetInt("DB_MAX_IDLE_CONNECTION")
	DB_MAX_OPEN_CONNECTION = viper.GetInt("DB_MAX_OPEN_CONNECTION")
	DB_CONN_MAX_LIFETIME_SECONDS = time.Second * time.Duration(viper.GetInt("DB_CONN_MAX_LIFETIME_SECONDS"))

	viper.SetDefault("CLOUDINARY_UPLOAD_FOLDER", "matchoshop")
	CLOUDINARY_URL = viper.GetString("CLOUDINARY_URL")
	CLOUDINARY_UPLOAD_FOLDER = viper.GetString("CLOUDINARY_UPLOAD_FOLDER")

	PAYPAL_CLIENT_ID = viper.GetString("PAYPAL_CLIENT_ID")
}

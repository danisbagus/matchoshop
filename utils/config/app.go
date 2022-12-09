package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	godotenv.Load()
	InitSqlite()
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	return port
}

func GetDBFile() string {
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "matchoshop.sqlite"
	}
	return dbFile
}

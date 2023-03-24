package modules

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetPostgresClient() *sqlx.DB {
	dbURL := os.Getenv("DATABASE_URL")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connection := fmt.Sprintf("%s?sslmode=%s", dbURL, dbSSLMode)

	client, err := sqlx.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}

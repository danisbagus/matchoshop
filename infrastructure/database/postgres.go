package database

import (
	"fmt"
	"log"
	"time"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host              string
	Port              int
	Username          string
	Password          string
	Database          string
	SSLMode           string
	MaxIdleConnection int
	MaxOpenConnection int
	MaxLiftime        time.Duration
}

func openPostgres(postgresConfig PostgresConfig) *sqlx.DB {
	connection := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", postgresConfig.Username, postgresConfig.Password, postgresConfig.Host, postgresConfig.Port, postgresConfig.Database, postgresConfig.SSLMode)
	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(postgresConfig.MaxLiftime)
	db.SetMaxOpenConns(postgresConfig.MaxOpenConnection)
	db.SetMaxIdleConns(postgresConfig.MaxIdleConnection)

	logger.Info("successfully connected to postgres database")
	return db
}

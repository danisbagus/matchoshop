package database

import (
	"github.com/danisbagus/matchoshop/infrastructure/config"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Postgres *PostgresConfig
}

type DatabaseConnection struct {
	Postgres *sqlx.DB
}

var (
	dbConnection DatabaseConnection
)

func GetConfigs() DatabaseConfig {
	postgresConfig := PostgresConfig{
		Host:              config.POSTGRES_HOST,
		Port:              config.POSTGRES_PORT,
		Username:          config.POSTGRES_USERNAME,
		Password:          config.POSTGRES_PASSWORD,
		Database:          config.POSTGRES_DATABASE,
		SSLMode:           config.POSTGRES_SSL_MODE,
		MaxIdleConnection: config.DB_MAX_IDLE_CONNECTION,
		MaxOpenConnection: config.DB_MAX_OPEN_CONNECTION,
		MaxLiftime:        config.DB_CONN_MAX_LIFETIME_SECONDS,
	}

	return DatabaseConfig{
		Postgres: &postgresConfig,
	}
}

// CreateDBConnections creates all database connections used in this app
//
// It will close the existing DB connection before opening a new DB connection
func CreateDBConnections(dbConfigs DatabaseConfig) DatabaseConnection {
	closeDBConnections(dbConnection)
	return openDBConnections(dbConfigs)

}

// openDBConnections opend database connection for each database
func openDBConnections(dbConfig DatabaseConfig) DatabaseConnection {
	var pgConnections *sqlx.DB

	if dbConfig.Postgres != nil {
		pgConnections = openPostgres(*dbConfig.Postgres)
	}

	dbConnection = DatabaseConnection{
		Postgres: pgConnections,
	}

	return dbConnection
}

// closeDBConnections close the existing connections if open
func closeDBConnections(dbConnection DatabaseConnection) {
	if dbConnection.Postgres != nil {
		dbConnection.Postgres.Close()
	}

	// close another DB (mysql, mongo, etc)
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/danisbagus/matchoshop/infrastructure/config"
	"github.com/danisbagus/matchoshop/infrastructure/database"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"

	//driver
	_ "github.com/lib/pq"
)

var usageCommands = `
Run database migrations & seeder
Usage:
    [command]
Available Migration Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version
`

func main() {
	godotenv.Load()
	var rootCmd = &cobra.Command{
		Use:   "migrate",
		Short: "MySql Migration Service",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}

			config.SetConfig(".", ".env")
			databaseConfig := database.GetConfigs()
			postgresConfig := databaseConfig.Postgres

			connection := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", postgresConfig.Username, postgresConfig.Password, postgresConfig.Host, postgresConfig.Port, postgresConfig.Database, postgresConfig.SSLMode)
			db, err := sql.Open("postgres", connection)
			if err != nil {
				panic(err)
			}

			appPath, _ := os.Getwd()
			dir := appPath + "/infrastructure/database/migrations"
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
			command := args[0]
			cmdArgs := args[1:]
			if err := goose.Run(command, db, dir, cmdArgs...); err != nil {
				log.Fatalf("goose run: %v", err)
			}
		},
	}

	rootCmd.SetUsageTemplate(usageCommands)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

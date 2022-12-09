package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/danisbagus/matchoshop/utils/config"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"

	//driver
	_ "github.com/mattn/go-sqlite3"
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
		Short: "sqlite Migration Service",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}

			file := config.GetDBFile()

			_, err := os.Stat(file)
			if os.IsNotExist(err) {
				fmt.Println("file not exits")
				var file, err = os.Create(file)
				if err != nil {
					fmt.Println("error create file")
					panic(err)
				}

				defer file.Close()
			}

			goose.SetDialect("sqlite3")
			db, err := sql.Open("sqlite3", file)
			if err != nil {
				log.Fatalf("failed connect to sqlite: %v", err)
			}

			appPath, _ := os.Getwd()
			dir := appPath + "/app/migration/migrations"
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

package main

import (
	"log"
	"os"

	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/cmd/server"
	"github.com/urfave/cli"
)

func main() {
	// Create a new CLI app
	app := cli.NewApp()
	app.Name = "Matchoshop"
	app.Usage = "Matchoshop API Service"
	app.Version = "1.0.0"

	// Define the CLI commands
	app.Commands = []cli.Command{
		{
			Name:  "http",
			Usage: "Run HTTP Server",
			Action: func(c *cli.Context) error {
				logger.Info("Starting http service...")
				server := server.NewServer()
				server.Http.Start()
				return nil

			},
		},
	}

	// Run the CLI app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

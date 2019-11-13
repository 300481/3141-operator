package main

import (
	"log"
	"os"

	"github.com/300481/cda/pkg/cmd/cda"
	"github.com/urfave/cli"
)

var (
	app = cli.NewApp()
)

func init() {
}

func info() {
	app.Name = "Continuous Deployment Agent"
	app.Usage = "Continuous Deployment Agent for Kubernetes, triggered by GitHub Webhooks."
	app.Version = "0.0.0"
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run in server mode",
			Action: func(c *cli.Context) error {
				cda.NewAgent().Serve()
				return nil
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

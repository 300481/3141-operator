package main

import (
	"log"
	"os"

	"github.com/300481/3141-operator/pkg/cmd/operator"
	"github.com/urfave/cli"
)

var (
	app = cli.NewApp()
)

func init() {
}

func info() {
	app.Name = "3141 Operator"
	app.Usage = "Kubernetes Operator running commands triggered by Git Webhooks."
	app.Author = "Dennis Riemenschneider"
	app.Version = "0.1.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Run in server mode",
			Action: func(c *cli.Context) {
				operator.NewOperator().Serve()
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

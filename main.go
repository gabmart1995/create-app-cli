package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"create-app-cli/commands/compress"
	"create-app-cli/commands/generate"
	"create-app-cli/commands/new"
	"create-app-cli/commands/serve"
)

var app = &cli.App{
	Name:        "create-web-app",
	Usage:       "Create a template web ready to work",
	Description: "Generator of structures for web static's",
	Version:     "0.1",
	Authors: []*cli.Author{
		{Name: "Gabriel Martinez", Email: "gabmart1995@gmail.com"},
		{Name: "Alfonso Martinez", Email: "martalf1987@gmail.com"},
	},
}

// constructor del modulo
func init() {
	app.Commands = []*cli.Command{
		&new.NewCommand,
		&generate.GenerateCommand,
		&serve.ServeCommand,
		&compress.CompressCommand,
	}
}

func main() {

	err := app.Run(os.Args)

	if err != nil {
		log.Fatalln(err)
	}
}

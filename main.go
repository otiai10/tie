package main

import (
	"log"
	"os"

	"github.com/otiai10/too/too"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "too"
	app.Usage = "too"
	app.Description = too.Description
	app.Action = too.MainAction
	app.Flags = []cli.Flag{too.FlagCommand}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}

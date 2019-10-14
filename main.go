package main

import (
	"log"
	"os"

	"github.com/otiai10/too/too"
)

func main() {
	app := too.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}

package main

import (
	"log"
	"os"

	"github.com/otiai10/tie/tie"
)

func main() {
	app := tie.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}

package tie

import (
	"os"

	"github.com/urfave/cli"
)

// NewApp ...
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "tie"
	app.Usage = "tie"
	app.Description = Description
	app.Action = func(ctx *cli.Context) {
		MainAction(ctx, os.Stdin)
	}
	app.Flags = []cli.Flag{FlagCommand, FlagFile}
	return app
}

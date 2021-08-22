package too

import (
	"os"

	"github.com/urfave/cli"
)

// NewApp ...
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "too"
	app.Usage = "too"
	app.Description = Description
	app.Action = func(ctx *cli.Context) {
		MainAction(ctx, os.Stdin)
	}
	app.Flags = []cli.Flag{FlagCommand, FlagFile}
	return app
}

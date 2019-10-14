package too

import "github.com/urfave/cli"

// NewApp ...
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "too"
	app.Usage = "too"
	app.Description = Description
	app.Action = MainAction
	app.Flags = []cli.Flag{FlagCommand}
	return app
}

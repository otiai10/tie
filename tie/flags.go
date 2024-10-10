package tie

import "github.com/urfave/cli"

// FlagCommand can accept commands to be mixed in one-line
var FlagCommand = cli.StringSliceFlag{
	Name:  "cmd",
	Usage: "commands to be mixed",
}

var FlagFile = cli.StringFlag{
	Name:  "f",
	Usage: "commads from a config file",
}

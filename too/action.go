package too

import (
	"github.com/urfave/cli"
)

// MainAction ...
func MainAction(ctx *cli.Context) error {

	builder := NewBuilder()

	if cmds := ctx.StringSlice("cmd"); len(cmds) != 0 {
		for _, cmdline := range cmds {
			if err := builder.Add(cmdline); err != nil {
				return err
			}
		}
	} else {
		if err := builder.Accept(); err != nil {
			return err
		}
	}

	commands, err := builder.Build()
	if err != nil {
		return err
	}
	for _, c := range commands {
		c.Introduction().Print(ctx.App.Writer)
	}

	if err := Exec(ctx.App.Writer, commands...); err != nil {
		return err
	}
	return nil

}

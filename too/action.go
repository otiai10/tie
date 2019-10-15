package too

import (
	"io"

	"github.com/urfave/cli"
)

// MainAction ...
func MainAction(ctx *cli.Context, stdin io.Reader) error {

	builder := NewBuilder(stdin)

	if cmds := ctx.StringSlice("cmd"); len(cmds) != 0 {
		for _, cmdline := range cmds {
			builder.Add(cmdline)
			// if err := builder.Add(cmdline); err != nil {
			// 	return err
			// }
		}
	} else {
		if err := builder.Accept(); err != nil {
			return err
		}
	}

	commands := builder.Build()
	for _, c := range commands {
		c.Introduction().Print(ctx.App.Writer)
	}

	if err := Exec(ctx.App.Writer, commands...); err != nil {
		return err
	}
	return nil

}

package too

import (
	"io"
	"os"

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
	} else if fpath := ctx.String("f"); fpath != "" {
		config, err := Parse(fpath)
		if err != nil {
			return err
		}
		for _, cmdline := range config.Commands {
			if err := builder.Add(cmdline.Run); err != nil {
				return err
			}
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

	err := Exec(ctx.App.Writer, commands...)
	if _, ok := err.(ErrorInterrupted); ok {
		os.Exit(130)
	}

	return err

}

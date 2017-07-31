package too

import "github.com/urfave/cli"

// MainAction ...
func MainAction(ctx *cli.Context) error {

	builder := NewBuilder()
	for {
		if err := builder.Accept(); err != nil {
			return err
		}
		break
	}
	commands, err := builder.Build()
	if err != nil {
		return err
	}
	if err := Exec(commands...); err != nil {
		return err
	}
	return nil
}

package too

import (
	"bytes"
	"flag"
	"testing"

	. "github.com/otiai10/mint"
	"github.com/urfave/cli"
)

func Test_MainAction(t *testing.T) {
	app := NewApp()
	set := flag.NewFlagSet("tootest", 0)
	set.Var(&cli.StringSlice{"echo foo", "echo baa"}, "cmd", "")
	buf := bytes.NewBuffer(nil)
	app.Writer = buf
	ctx := cli.NewContext(app, set, nil)
	err := MainAction(ctx)
	Expect(t, err).ToBe(nil)
	Expect(t, buf.String()).ToBe("[0] echo\tfoo\n[1] echo\tbaa\n")
}

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
	Expect(t, buf.String()).Match("\\[0\\] echo foo")
	Expect(t, buf.String()).Match("\\[1\\] echo baa")
	Expect(t, buf.String()).Match("\\[0\\] echo\tfoo\n")
	Expect(t, buf.String()).Match("\\[1\\] echo\tbaa\n")
	Expect(t, buf.String()).Match("\\[0\\] echo\texit code 0\n")
	Expect(t, buf.String()).Match("\\[1\\] echo\texit code 0\n")

	When(t, "Some command exits with non-zero code", func(t *testing.T) {
		set := flag.NewFlagSet("tootest", 0)
		set.Var(&cli.StringSlice{"echo foo", "cat non-existing"}, "cmd", "")
		buf := bytes.NewBuffer(nil)
		app.Writer = buf
		ctx := cli.NewContext(app, set, nil)
		err := MainAction(ctx)
		Expect(t, err).ToBe(nil)
		Expect(t, buf.String()).Match("\\[0\\] echo foo")
		Expect(t, buf.String()).Match("\\[1\\] cat non-existing")
		Expect(t, buf.String()).Match("\\[1\\] cat\texit code 1\n")
	})
}

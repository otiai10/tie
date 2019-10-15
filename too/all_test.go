package too

import (
	"bufio"
	"bytes"
	"flag"
	"syscall"
	"testing"
	"time"

	"github.com/otiai10/color"
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
	err := MainAction(ctx, nil)
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
		err := MainAction(ctx, nil)
		Expect(t, err).ToBe(nil)
		Expect(t, buf.String()).Match("\\[0\\] echo foo")
		Expect(t, buf.String()).Match("\\[1\\] cat non-existing")
		Expect(t, buf.String()).Match("\\[1\\] cat\texit code 1\n")
	})

	When(t, "Any of cmd arguments are NOT given", func(t *testing.T) {
		app := NewApp()
		set := flag.NewFlagSet("tootest", 0)
		// set.Var(nil, "cmd", "")
		inbuf := bytes.NewBuffer([]byte("echo hoge\necho fuga\n\n"))
		outbuf := bytes.NewBuffer(nil)
		app.Writer = outbuf
		ctx := cli.NewContext(app, set, nil)
		err := MainAction(ctx, inbuf)
		Expect(t, err).ToBe(nil)
		Expect(t, outbuf.String()).Match("\\[0\\] echo hoge")
		Expect(t, outbuf.String()).Match("\\[1\\] echo fuga")

		Because(t, "Input is too long", func(t *testing.T) {
			instr := "echo hoge"
			for len(instr) < bufio.MaxScanTokenSize {
				instr = instr + instr
			}
			inbuf := bytes.NewBuffer([]byte(instr + "\n\n"))
			outbuf := bytes.NewBuffer(nil)
			app.Writer = outbuf
			ctx := cli.NewContext(app, set, nil)
			err := MainAction(ctx, inbuf)
			Expect(t, err).Not().ToBe(nil)
		})
	})

}

func TestBuilder_Accept(t *testing.T) {
	buf := bytes.NewBuffer([]byte("hoge\nfuga\npiyo\n\n"))
	builder := NewBuilder(buf)
	err := builder.Accept()
	Expect(t, err).ToBe(nil)
	commands := builder.Build()
	// Expect(t, err).ToBe(nil)
	Expect(t, len(commands)).ToBe(3)
}

func TestExec(t *testing.T) {
	output := bytes.NewBuffer(nil)
	cmd1 := NewCommand("echo foobaa", 0, color.New(color.FgCyan))
	cmd2 := NewCommand("sleep 10", 1, color.New(color.FgGreen))
	res := make(chan error)
	go func() {
		err := Exec(output, cmd1, cmd2)
		time.Sleep(500 * time.Millisecond)
		res <- err
	}()
	go func() {
		for {
			if cmd2.Cmd.Process != nil {
				break
			}
		}
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()
	err := <-res
	Expect(t, err).ToBe(nil)
	Expect(t, output.String()).Match("\\[0\\] echo	exit code 0")
	Expect(t, output.String()).Match("\\[1\\] sleep	exit code -1")
}

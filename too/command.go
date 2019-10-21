package too

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"

	"github.com/otiai10/color"
	"github.com/otiai10/spell"
)

// Command ...
type Command struct {
	stdout   io.Writer
	Color    *color.Color
	Index    int
	Prefix   string
	RawInput string
	*exec.Cmd
}

// NewCommand ...
func NewCommand(line string, index int, col *color.Color) *Command {
	c := &Command{
		Index:    index,
		Color:    col,
		RawInput: line,
	}

	envexp := regexp.MustCompile("^[^=]+=[^=]+$")
	var envs, spells []string
	for _, token := range spell.Parse(line) {
		if envexp.MatchString(token) {
			envs = append(envs, token)
		} else {
			spells = append(spells, token)
		}
	}

	c.Prefix = spells[0]

	if len(spells) > 1 {
		c.Cmd = exec.Command(spells[0], spells[1:]...)
	} else {
		c.Cmd = exec.Command(spells[0])
	}
	c.Env = append(os.Environ(), envs...)
	return c
}

// Start ...
func (c *Command) Start(msg chan<- Message, end chan *Command) error {

	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}

	go func() {
		scanout := bufio.NewScanner(stdout)
		for scanout.Scan() {
			msg <- c.message(Stdout, scanout.Text())
		}
		stdout.Close()
		end <- c
	}()

	go func() {
		scanerr := bufio.NewScanner(stderr)
		for scanerr.Scan() {
			msg <- c.message(Stderr, scanerr.Text())
		}
		stderr.Close()
	}()

	if err := c.Cmd.Start(); err != nil {
		stdout.Close()
		stderr.Close()
		return err
	}

	return nil

}

// Introduction prints raw input with underline.
func (c *Command) Introduction() Message {
	withDecoration := c.Color.Clone()
	withDecoration.Add(color.Underline)
	return Message{
		Output: Stderr,
		Color:  withDecoration,
		Header: fmt.Sprintf("[%d] %s", c.Index, c.RawInput),
	}
}

// ExitCode ...
func (c *Command) ExitCode() Message {
	return c.message(Stderr, fmt.Sprintf("exit code %d", c.Cmd.ProcessState.ExitCode()))
}

func (c *Command) message(out Output, text string) Message {
	return Message{
		Output: out,
		Color:  c.Color,
		Header: fmt.Sprintf("[%d] %s\t", c.Index, c.Prefix),
		Text:   text,
	}
}

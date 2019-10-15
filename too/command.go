package too

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/otiai10/color"
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
func NewCommand(line string, index int, col *color.Color) (*Command, error) {
	c := &Command{
		Index:    index,
		Color:    col,
		RawInput: line,
	}
	q := strings.Split(line, " ")
	if len(q) == 0 {
		return nil, fmt.Errorf("invalid line")
	}
	envs, spell := parseWords(q)
	c.Prefix = spell[0]
	if len(q) > 1 {
		c.Cmd = exec.Command(spell[0], spell[1:]...)
	} else {
		c.Cmd = exec.Command(spell[0])
	}
	c.Env = append(os.Environ(), envs...)
	return c, nil
}

func parseWords(words []string) ([]string, []string) {
	exp := regexp.MustCompile("^[^=]+=[^=]+$")
	envs := []string{}
	spell := []string{}
	for _, word := range words {
		if exp.MatchString(word) {
			envs = append(envs, word)
		} else {
			spell = append(spell, word)
		}
	}
	return envs, spell
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

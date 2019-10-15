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
func NewCommand(stdout io.Writer, line string, index int, col *color.Color) (*Command, error) {
	c := &Command{
		stdout:   stdout,
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

// PrintIntroduction prints raw input with underline.
func (c *Command) PrintIntroduction() {
	out := c.stdout
	if out == nil {
		out = os.Stdout
	}
	withDecoration := c.Color.Clone()
	withDecoration.Add(color.Underline)
	withDecoration.Fprintf(out, "[%d] %s\n", c.Index, c.RawInput)
}

// PrintHeader ...
func (c *Command) PrintHeader() {
	out := c.stdout
	if out == nil {
		out = os.Stdout
	}
	c.Color.Fprintf(out, "[%d] %s\t", c.Index, c.Prefix)
}

// PrintLine ...
func (c *Command) PrintLine(text string) {
	out := c.stdout
	if out == nil {
		out = os.Stderr
	}
	fmt.Fprintln(out, text)
	// c.Color.Fprintln(out, text)
}

// PrintExitCode ...
func (c *Command) PrintExitCode() {
	out := c.stdout
	if out == nil {
		out = os.Stdout
	}
	c.Color.Fprintf(out, "[%d] %s\t", c.Index, c.Prefix)
	fmt.Fprintf(out, "exit code %d\n", c.Cmd.ProcessState.ExitCode())
}

// Start ...
func (c *Command) Start(end chan *Command) error {

	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}

	scout := bufio.NewScanner(stdout)
	go func() {
		for scout.Scan() {
			c.PrintHeader()
			c.PrintLine(scout.Text())
		}
		stdout.Close()
		end <- c
	}()

	scerr := bufio.NewScanner(stderr)
	go func() {
		for scerr.Scan() {
			c.PrintHeader()
			c.PrintLine(scerr.Text())
		}
		stderr.Close()
		// end <- true
	}()

	c.PrintIntroduction()

	if err := c.Cmd.Start(); err != nil {
		stdout.Close()
		stderr.Close()
		return err
	}

	return nil

}

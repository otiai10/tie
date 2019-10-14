package too

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

// Command ...
type Command struct {
	stdout io.Writer
	Color  *color.Color
	Index  int
	Prefix string
	*exec.Cmd
}

// NewCommand ...
func NewCommand(stdout io.Writer, line string, index int, col *color.Color) (*Command, error) {
	c := &Command{
		stdout: stdout,
		Index:  index,
		Color:  col,
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
	c.Color.Fprintln(out, text)
}

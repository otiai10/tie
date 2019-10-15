package too

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/otiai10/color"
)

var colors = []color.Attribute{
	color.FgGreen,
	color.FgCyan,
}

const prompt = "> "

// Builder ...
type Builder struct {
	commands []*Command
	stdout   io.Writer
}

// NewBuilder ...
func NewBuilder(writer io.Writer) *Builder {
	return &Builder{
		stdout: writer,
	}
}

// Accept ...
func (b *Builder) Accept() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if line == "" {
			return nil
		}
		if err := b.Add(line); err != nil {
			return err
		}
		fmt.Print(prompt)
	}
	return nil
}

// Add ...
func (b *Builder) Add(line string) error {
	cmd, err := NewCommand(b.stdout, line, len(b.commands), color.New(colors[len(b.commands)%len(colors)]))
	if err != nil {
		return err
	}
	b.commands = append(b.commands, cmd)
	return nil
}

// Build ...
func (b *Builder) Build() ([]*Command, error) {
	return b.commands, nil
}

package too

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/otiai10/color"
)

var colors = []color.Attribute{
	color.FgGreen,
	color.FgCyan,
	color.FgHiYellow,
	color.FgMagenta,
}

const prompt = "> "

// Builder ...
type Builder struct {
	commands []*Command
	stdin    io.Reader
}

// NewBuilder ...
func NewBuilder(stdin io.Reader) *Builder {
	return &Builder{
		stdin: stdin,
	}
}

// Accept ...
func (b *Builder) Accept() error {
	scanner := bufio.NewScanner(b.stdin)
	fmt.Print(prompt)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if line == "" {
			break
		}
		b.Add(line)
		fmt.Print(prompt)
	}
	return scanner.Err()
}

// Add ...
func (b *Builder) Add(line string) error {
	cmd := NewCommand(line, len(b.commands), color.New(colors[len(b.commands)%len(colors)]))
	b.commands = append(b.commands, cmd)
	return nil
}

// Build ...
func (b *Builder) Build() []*Command {
	// return b.commands, nil
	return b.commands
}

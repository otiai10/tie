package too

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// Command ...
type Command struct {
	Color *color.Color
	Index int
	*exec.Cmd
}

// NewCommand ...
func NewCommand(line string, index int, col *color.Color) (*Command, error) {
	c := &Command{
		Index: index,
		Color: col,
	}
	q := strings.Split(line, " ")
	if len(q) == 0 {
		return nil, fmt.Errorf("invalid line")
	}
	if len(q) > 1 {
		c.Cmd = exec.Command(q[0], q[1:]...)
	} else {
		c.Cmd = exec.Command(q[0])
	}
	return c, nil
}

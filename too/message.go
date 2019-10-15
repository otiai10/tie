package too

import (
	"fmt"
	"io"

	"github.com/otiai10/color"
)

// Output ...
type Output int

const (
	// Stdout ...
	Stdout Output = iota + 1
	// Stderr ...
	Stderr

	// AppEnd ...
	AppEnd = -1
)

// Message ...
type Message struct {
	Output Output
	Color  *color.Color
	Header string
	Text   string
}

// Print ...
func (m Message) Print(out io.Writer) {
	m.Color.Fprintf(out, "%s", m.Header)
	fmt.Fprintf(out, "\t%s\n", m.Text)
}

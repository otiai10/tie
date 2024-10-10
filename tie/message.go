package tie

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
)

// Message ...
type Message struct {
	Output Output
	Color  *color.Color
	Header string
	Text   string
}

// AppEnd ...
var AppEnd = Message{Output: -1}

// Print ...
func (m Message) Print(out io.Writer) {
	m.Color.Fprintf(out, "%s", m.Header)
	fmt.Fprintf(out, "\t%s\n", m.Text)
}

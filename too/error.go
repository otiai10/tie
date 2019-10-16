package too

import "fmt"

// ErrorInterrupted ...
type ErrorInterrupted struct {
	CommandErrors []error
}

// Error to satisfy error interface.
func (e ErrorInterrupted) Error() string {
	msg := "commands interrupted"
	if len(e.CommandErrors) == 0 {
		return msg
	}
	msg += " with errors:\n"
	for _, err := range e.CommandErrors {
		msg += "\t" + err.Error() + "\n"
	}
	return msg + "\n"
}

// Add command specific errors.
func (e ErrorInterrupted) Add(bin string, err error) {
	e.CommandErrors = append(e.CommandErrors, fmt.Errorf("%s: %v", bin, err))
}

package too

import (
	"fmt"
	"io"
	"os"
	"os/signal"
)

// Exec ...
func Exec(output io.Writer, commands ...*Command) error {

	interrupt := make(chan os.Signal, 1)
	// defer close(interrupt)
	signal.Notify(interrupt, os.Interrupt)

	endups := make(chan *Command)
	defer close(endups)
	endcnt := 0

	stdout := output
	stderr := output
	if stderr == os.Stdout {
		stderr = os.Stderr
	}
	msg := make(chan Message)
	go msgQueue(stdout, stderr, msg)

	for _, c := range commands {
		if err := c.Start(msg, endups); err != nil {
			return err
		}
	}

	for {
		select {
		case _ = <-interrupt:
			errors := []error{}
			for _, c := range commands {
				if c.ProcessState != nil && c.ProcessState.Exited() {
					// This process is already released
					continue
				}
				if err := c.Process.Kill(); err != nil {
					errors = append(errors, err)
				}
			}
			if len(errors) != 0 {
				fmt.Println(errors)
			}
			break
		case c := <-endups:
			endcnt++
			c.Wait()
			msg <- c.ExitCode()
			if endcnt >= len(commands) {
				// To make sure to print the last message
				msg <- Message{Output: AppEnd}
				return nil
			}
		}
	}
}

func msgQueue(stdout, stderr io.Writer, msgchan chan Message) {
	for {
		m := <-msgchan
		switch m.Output {
		case AppEnd:
			close(msgchan)
			return
		case Stdout:
			m.Color.Fprintf(stdout, m.Header)
			fmt.Fprintf(stdout, "%s\n", m.Text)
		case Stderr:
			m.Color.Fprintf(stderr, m.Header)
			fmt.Fprintf(stderr, "%s\n", m.Text)
		}
	}
}

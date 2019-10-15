package too

import (
	"fmt"
	"os"
	"os/signal"
)

// Exec ...
func Exec(commands ...*Command) error {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	endups := make(chan *Command)
	endcnt := 0

	for _, c := range commands {
		if err := c.Start(endups); err != nil {
			return err
		}
	}

	for {
		select {
		case _ = <-interrupt:
			errors := []error{}
			for _, c := range commands {
				if err := c.Process.Kill(); err != nil {
					errors = append(errors, err)
				}
			}
			if len(errors) != 0 {
				fmt.Println(errors)
			}
			// os.Exit(0)
			break
		case c := <-endups:
			endcnt++
			c.Wait()
			c.PrintExitCode()
			if endcnt >= len(commands) {
				return nil
			}
		}
	}
}

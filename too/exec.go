package too

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

// start ...
func start(c *Command, end chan bool) error {

	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}

	output := c.stdout
	if output == nil {
		output = os.Stdout
	}

	scout := bufio.NewScanner(stdout)
	go func() {
		for scout.Scan() {
			c.PrintHeader()
			c.PrintLine(scout.Text())
		}
		stdout.Close()
		end <- true
	}()

	scerr := bufio.NewScanner(stderr)
	go func() {
		for scerr.Scan() {
			c.PrintHeader()
			c.PrintLine(scout.Text())
		}
		stderr.Close()
		// end <- true
	}()

	if err := c.Start(); err != nil {
		stdout.Close()
		stderr.Close()
		return err
	}

	return nil
}

// Exec ...
func Exec(commands ...*Command) error {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	endups := make(chan bool)
	endcnt := 0

	for _, c := range commands {
		if err := start(c, endups); err != nil {
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
		case _ = <-endups:
			endcnt++
			if endcnt >= len(commands) {
				// os.Exit(0)
				return nil
			}
		}
	}
}

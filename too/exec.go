package too

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

// start 1 command and notify when it ends by the channel provided.
func start(c *Command, end chan *Command) error {

	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}

	scout := bufio.NewScanner(stdout)
	go func() {
		for scout.Scan() {
			c.PrintHeader()
			c.PrintLine(scout.Text())
		}
		stdout.Close()
		end <- c
	}()

	scerr := bufio.NewScanner(stderr)
	go func() {
		for scerr.Scan() {
			c.PrintHeader()
			c.PrintLine(scerr.Text())
		}
		stderr.Close()
		// end <- true
	}()

	c.PrintIntroduction()

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

	endups := make(chan *Command)
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

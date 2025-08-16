package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/This-Is-Prince/goshelly/app/cmd"
)

func main() {
	c := cmd.NewCmd("")
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		c.Reset(command)
		c.Evaluate()
	}
}

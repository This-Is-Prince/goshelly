package cmd

import (
	"fmt"
)

func (c *Cmd) Evaluate() {
	c.BuildRunnableCmd()

	var msg string
	var err error

	run, isBuiltin := BUILT_IN_CMDS[c.rCmd]

	if isBuiltin {
		msg, err = run(c.rCmd, c.rArgsCmd)
	} else {
		fmt.Fprintf(c.Stdout, "%s: command not found", c.rCmd)
		return
	}

	if err != nil {
		fmt.Fprintf(c.Stderr, "%v\n", err)
	} else if msg != "" {
		fmt.Fprintf(c.Stdout, "%s", msg)
	}
}

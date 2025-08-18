package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Cmd struct {
	raw        string
	kind       string
	cleanedCmd string
	args       []string

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func NewCmd(raw string) *Cmd {
	return &Cmd{
		raw:    raw,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

func (c *Cmd) Clean() *Cmd {
	c.cleanedCmd = strings.TrimSpace(c.raw)

	return c
}

func (c *Cmd) Build() *Cmd {
	separator := " "

	splitCmd := strings.Split(c.cleanedCmd, separator)

	if len(splitCmd) > 0 {
		c.kind = splitCmd[0]
		c.args = splitCmd[1:]
	}

	return c
}

func (c *Cmd) Run() (output string, err error) {
	run, isBuiltin := BUILT_IN_CMDS[c.kind]

	if isBuiltin {
		output, err = run(c.kind, c.args)
	} else {
		output = fmt.Sprintf("%s: command not found\n", c.kind)
	}

	return
}

func (c *Cmd) Reset(raw string) {
	c.raw = raw
	c.kind = ""
	c.cleanedCmd = ""
	c.args = []string{}

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
}

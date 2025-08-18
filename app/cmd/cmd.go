package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
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
		pathEnv := os.Getenv("PATH")
		separator := string(os.PathListSeparator)

		isExecutable := false

		for _, path := range strings.Split(pathEnv, separator) {
			path = filepath.Join(path, c.kind)
			if info, err := os.Stat(path); err == nil {
				mode := info.Mode()

				if mode.Perm()&0111 != 0 {
					isExecutable = true
					break
				}
			}
		}

		if !isExecutable {
			output = fmt.Sprintf("%s: command not found\n", c.kind)
			return
		}

		eCmd := exec.Command(c.kind, c.args...)

		var outputPipe, errorPipe io.ReadCloser

		outputPipe, err = eCmd.StdoutPipe()
		if err != nil {
			return
		}

		errorPipe, err = eCmd.StderrPipe()
		if err != nil {
			return
		}

		err = eCmd.Start()
		if err != nil {
			return
		}

		var outputBytes, errorBytes []byte
		outputBytes, _ = io.ReadAll(outputPipe)
		errorBytes, _ = io.ReadAll(errorPipe)

		err = eCmd.Wait()
		if err != nil {
			return
		}

		if len(outputBytes) > 0 {
			output = fmt.Sprint(string(outputBytes))
		} else if len(errorBytes) > 0 {
			err = fmt.Errorf("%v", string(errorBytes))
		}
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

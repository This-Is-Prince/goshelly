package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/This-Is-Prince/goshelly/app/cmd"
)

type Shell struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	err     error
	output  string
	command string
	cmd     *cmd.Cmd
}

func NewShell() (shell *Shell) {
	shell = &Shell{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		cmd:    cmd.NewCmd(""),
	}

	return
}

func (s *Shell) REPL() {
	for {
		s.Read()
		s.Evaluate()
		s.Print()
		s.Loop()
	}
}

func (s *Shell) Read() {
	fmt.Fprint(s.Stdout, "$ ")

	command, err := bufio.NewReader(s.Stdin).ReadString('\n')
	if err != nil {
		s.err = err
		return
	}

	s.command = command
}

func (s *Shell) Evaluate() {
	if s.err != nil {
		return
	}

	c := s.cmd

	c.Reset(s.command)
	output, err := c.Clean().Build().Run()

	s.err = err
	s.output = output
}

func (s *Shell) Print() {
	if s.err != nil {
		fmt.Fprintf(s.Stderr, "%v\n", s.err)
	} else if s.output != "" {
		fmt.Fprintf(s.Stdout, "%v", s.output)
	}
}

func (s *Shell) Loop() {
	s.err = nil
	s.output = ""
	s.command = ""
}

package cmd

import "strings"

type Cmd struct {
	rawCmd     string
	cleanedCmd string
	rCmd       string
	rArgsCmd   []string
}

func NewCmd(rawCmd string) *Cmd {
	return &Cmd{
		rawCmd: rawCmd,
	}
}

func (cmd *Cmd) Reset(rawCmd string) {
	cmd.rawCmd = rawCmd
	// All other reset operations
}

func (cmd *Cmd) BuildRunnableCmd() {
	cmd.Clean()

	separator := " "

	splitCmd := strings.Split(cmd.cleanedCmd, separator)

	if len(splitCmd) > 0 {
		cmd.rCmd = splitCmd[0]
		cmd.rArgsCmd = splitCmd[1:]
	}
}

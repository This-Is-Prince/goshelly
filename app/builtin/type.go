package builtin

import "fmt"

type emptyType struct{}

var BUILT_IN_CMDS = map[string]struct{}{
	"exit": emptyType{},
	"echo": emptyType{},
	"type": emptyType{},
}

func Type(cmd string, args []string) (output string, err error) {
	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		_, ok := BUILT_IN_CMDS[arg]
		if ok {
			output += fmt.Sprintf("%s is a shell builtin\n", arg)
		} else {
			output += fmt.Sprintf("%s: not found\n", arg)
		}
	}

	return
}

package builtin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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

	pathEnv := os.Getenv("PATH")

args:
	for _, arg := range args {
		_, ok := BUILT_IN_CMDS[arg]
		if ok {
			output += fmt.Sprintf("%s is a shell builtin\n", arg)
		} else {
			separator := ":"

			for path := range strings.SplitSeq(pathEnv, separator) {
				path = filepath.Join(path, arg)
				if info, err := os.Stat(path); err == nil {
					mode := info.Mode()

					if mode.Perm()&0111 != 0 {
						output += fmt.Sprintf("%s is %v\n", arg, path)
						continue args
					}
				}
			}

			output += fmt.Sprintf("%s: not found\n", arg)
		}
	}

	return
}

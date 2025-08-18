package builtin

import (
	"fmt"
	"os"
)

func Cd(cmd string, args []string) (output string, err error) {
	if len(args) > 1 {
		err = fmt.Errorf("%s: too many arguments", cmd)
		return
	}

	path := args[0]

	if path == "~" {
		path, err = os.UserHomeDir()
		if err != nil {
			return
		}
	}

	err = os.Chdir(path)

	return
}

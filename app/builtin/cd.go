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

	var target string

	if len(args) == 0 || args[0] == "" || args[0] == "~" {
		if target, err = os.UserHomeDir(); err != nil {
			return
		}
	} else {
		target = args[0]
	}

	err = os.Chdir(target)

	return
}

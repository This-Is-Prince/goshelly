package builtin

import (
	"fmt"
	"os"
)

func Pwd(cmd string, args []string) (output string, err error) {
	if len(args) > 0 {
		err = fmt.Errorf("%s: too many arguments", cmd)
		return
	}

	output, err = os.Getwd()

	if err == nil {
		output = fmt.Sprintf("%s\n", output)
	}

	return
}

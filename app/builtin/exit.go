package builtin

import (
	"fmt"
	"os"
)

var ErrTooManyArguments = fmt.Errorf("exit: too many arguments")
var ErrInvalidExitCode = fmt.Errorf("exit: invalid exit code")

func Exit(cmd string, args []string) (string, error) {
	if len(args) > 1 {
		return "", ErrTooManyArguments
	}

	if args[0] != "0" {
		return "", ErrInvalidExitCode
	}

	os.Exit(0)

	return "", nil
}

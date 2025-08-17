package builtin

import (
	"fmt"
	"os"
	"strconv"
)

var (
	ErrTooManyArguments = fmt.Errorf("exit: too many arguments")
	ErrInvalidExitCode  = fmt.Errorf("exit: invalid exit code")
)

var exitFn = os.Exit

// Exit implements: exit [n],
// no args -> exit 0,
// one numeric arg 0..255 -> exit with that code,
// invalid/too many args -> error (do not exit),
func Exit(cmd string, args []string) (string, error) {
	if len(args) > 1 {
		return "", ErrTooManyArguments
	}

	code := 0
	if len(args) == 1 {
		n, err := strconv.Atoi(args[0])
		if err != nil || n < 0 || n > 255 {
			return "", ErrInvalidExitCode
		}
		code = n
	}

	exitFn(code)

	return "", nil
}

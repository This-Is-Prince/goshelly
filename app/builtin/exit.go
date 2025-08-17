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
func Exit(cmd string, args []string) (msg string, err error) {
	if len(args) > 1 {
		err = ErrTooManyArguments
		return
	}

	code := 0
	if len(args) == 1 {
		n, e := strconv.Atoi(args[0])
		if e != nil || n < 0 || n > 255 {
			err = ErrInvalidExitCode
			return
		}
		code = n
	}

	exitFn(code)

	msg = "\n"

	return
}

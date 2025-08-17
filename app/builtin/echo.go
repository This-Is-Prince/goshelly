package builtin

import (
	"strings"
)

func Echo(cmd string, args []string) (msg string, err error) {
	_args := []string{}

	for _, arg := range args {
		arg := strings.TrimSpace(arg)
		if arg != "" {
			_args = append(_args, arg)
		}
	}

	if len(_args) > 0 {
		msg = strings.Join(_args, " ")
	}

	msg += "\n"

	return
}

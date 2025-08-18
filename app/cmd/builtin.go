package cmd

import "github.com/This-Is-Prince/goshelly/app/builtin"

var BUILT_IN_CMDS = map[string]func(cmd string, args []string) (string, error){
	"exit": builtin.Exit,
	"echo": builtin.Echo,
	"type": builtin.Type,
	"pwd":  builtin.Pwd,
	"cd":   builtin.Cd,
}

package cmd

import "strings"

func (c *Cmd) Clean() {
	c.cleanedCmd = strings.TrimSpace(c.rawCmd)
}

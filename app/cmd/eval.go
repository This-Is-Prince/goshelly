package cmd

func (c *Cmd) Evaluate() {
	c.BuildRunnableCmd()

	run, ok := BUILT_IN_CMDS[c.rCmd]

	if ok {
		run(c.rCmd, c.rArgsCmd)
	}
}

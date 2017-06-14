package main

import (
	"github.com/mitchellh/cli"
)

type Finish struct {
	ui cli.Ui
}

func (f *Finish) Help() string {
	return "task-manage finish <task name>"
}

func (f *Finish) Run(args []string) int {
	db.Finish(args[0])
	return 0
}

func (f *Finish) Synopsis() string {
	return "done Task"
}

package main

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type Delete struct {
	ui cli.Ui
}

func (f *Delete) Help() string {
	return "task-manage delete <-h have, -f finished>"
}

func (f *Delete) Run(args []string) int {
	stateFlag := args[1]
	state := parseStateFlag(stateFlag)
	title := args[0]

	if err := db.deleteOf(title, state); err != nil {
		f.ui.Error(fmt.Sprint(err))
		return 1
	}
	f.ui.Info(fmt.Sprintf("Deleted: %s", args[0]))
	return 0
}

func (f *Delete) Synopsis() string {
	return "display of all tasks"
}

func parseStateFlag(s string) TaskState {
	if s == "-o" {
		return ongoing
	}
	return finished
}

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

type Done struct {
	ui cli.Ui
}

func (f *Done) Help() string {
	return "task-manage done <task name>"
}

func (f *Done) Run(args []string) int {
	old_file := filepath.Join(root.have, args[0]+".json")
	data := FindTask(old_file)
	data.DoneTime, _ = NewPoint(0)
	creatore.Task(data, root.have)
	new_file := filepath.Join(root.finished, args[0]+".json")
	err := os.Rename(old_file, new_file)
	if err != nil {
		f.ui.Error(fmt.Sprint(err))
		return 1
	}
	return 0
}

func (f *Done) Synopsis() string {
	return "done Task"
}

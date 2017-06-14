package main

import (
	"os"

	"github.com/mitchellh/cli"
)

type Show struct {
	ui cli.Ui
}

func (f *Show) Help() string {
	return "app Show"
}

func (f *Show) Run(args []string) int {
	if len(os.Args) < 3 {
		paths := db.Ongoing()
		for _, path := range paths {
			task, err := db.readFrom(path)
			if err != nil {
				f.ui.Error(err.Error())
				return 1
			}
			printTask(task)
		}
	} else {
		task, err := db.readFrom(db.calcFileName(args[0]))
		if err != nil {
			f.ui.Error(err.Error())
			return 1
		}
		printTask(task)
	}
	return 0
}

func (f *Show) Synopsis() string {
	return "display of task or all task"
}

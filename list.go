package main

import "github.com/mitchellh/cli"

type List struct {
	ui cli.Ui
}

func (f *List) Help() string {
	return "app List"
}

func (f *List) Run(args []string) int {
	paths := db.All()
	for _, path := range paths {
		task, err := db.readFrom(path)
		if err != nil {
			f.ui.Error(err.Error())
			return 1
		}
		printTask(task)
	}
	return 0
}

func (f *List) Synopsis() string {
	return "display of all tasks"
}

package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mitchellh/cli"
)

type Add struct {
	ui cli.Ui
}

func (f *Add) Help() string {
	return "You can add a task \n task-manage add <title> <content> [<Days>]"
}

func (f *Add) Run(args []string) int {
	n, err := strconv.Atoi(args[2])
	if err != nil {
		f.ui.Error(fmt.Sprint(err))
		return 1
	}
	task := Task{
		Title: args[0],
		Content: args[1],
		CreatedAt: time.Now(),
		Deadline: time.Now().AddDate(0, 0, n),
	}

	err = db.Store(task)
	if err != nil {
		f.ui.Error(fmt.Sprintf("%s: %s\n", task.Title, err))
		return 1
	}

	return 0
}

func (f *Add) Synopsis() string {
	return "add Task"
}

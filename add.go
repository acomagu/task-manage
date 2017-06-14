package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mitchellh/cli"
)

type Data struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	DeadLine  time.Time `json:"deadline"`
	DoneTime  time.Time `json:"donetime"`
}

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
	end, now := NewPoint(n)
	data := Task{
		Title: args[0],
		Content: args[1],
		CreatedAt: now,
		Deadline: end,
	}

	err = creatore.Task(data, root.have)
	if err != nil {
		f.ui.Error(fmt.Sprintf("%s: %s", data.Title, err))
		return 1
	}

	return 0
}

func (f *Add) Synopsis() string {
	return "add Task"
}

package main

import (
	"strconv"
)

type Add struct{}

func (f *Add) Help() string {
	return "You can add a task \n task-manage add <title> <content> [<Days>]"
}

func (f *Add) Run(args []string) int {
	n, _ := strconv.Atoi(args[2])
	end, now := NewPoint(n)
	data := Task{
		Title: args[0],
		Content: args[1],
		CreatedAt: now,
		Deadline: end,
	}
	creatore.Task(data, root.have)
	return 0
}

func (f *Add) Synopsis() string {
	return "add Task"
}

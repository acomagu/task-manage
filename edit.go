package main

import (
	"fmt"
	"os"
	"strconv"
)

type Edit struct{}

func (f *Edit) Help() string {
	return "task-manage edit <task Title> <task New Title> <task New Content> <task Nwe days>"
}

func (f *Edit) Run(args []string) int {
	task := root.have + args[0] + ".json"
	data := FindTask(task)
	if err := os.Remove(task); err != nil {
		fmt.Println(err)
	}
	data.Title = args[1]
	data.Content = args[2]
	n, _ := strconv.Atoi(args[3])
	end, _ := NewPoint(n)
	data.DeadLine = end
	creatore.Task(data, root.have)
	return 0
}

func (f *Edit) Synopsis() string {
	return "edit task"
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Edit struct{}

func (f *Edit) Help() string {
	return "task-manage edit <task Title> <task New Title> <task New Content> <task Nwe days>"
}

func (f *Edit) Run(args []string) int {
	task := filepath.Join(root.have, args[0]+".json")
	data := FindTask(task)
	if err := os.Remove(task); err != nil {
		fmt.Println(err)
	}
	data.Title = args[1]
	data.Content = args[2]
	n, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println(err)
	}
	end, _ := NewPoint(n)
	data.DeadLine = end
	data.DoneTime = end
	creatore.Task(data, root.have)
	return 0
}

func (f *Edit) Synopsis() string {
	return "edit task"
}

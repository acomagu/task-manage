package main

import (
	"fmt"
	"os"
)

type Done struct{}

func (f *Done) Help() string {
	return "task-manage done <task name>"
}

func (f *Done) Run(args []string) int {
	old_file := root.have + os.Args[2] + ".json"
	data := FindTask(old_file)
	data.DoneTime, _ = NewPoint(0)
	creatore.Task(data, root.have)
	new_file := root.finished + os.Args[2] + ".json"
	err := os.Rename(old_file, new_file)
	if err != nil {
		fmt.Println(err)
	}
	return 0
}

func (f *Done) Synopsis() string {
	return "done Task"
}

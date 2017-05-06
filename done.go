package main

import (
  "os"
  "fmt"
)

type Done struct{}

func (f *Done) Help() string {
	return "task-manage done <task name>"
}

func (f *Done) Run(args []string) int {
  path := os.Getenv("GOPATH")
  root := path + "/src/github.com/Yamashou/task-manage/Tasks/Have/"
  old_file := root + os.Args[2] + "_Task.json"
  new_file := path +"/src/github.com/Yamashou/task-manage/Tasks/Finished/" + os.Args[2] + "_Done.json"

  err := os.Rename(old_file, new_file)
  if err != nil {
		fmt.Println(1, err)
	}
	return 0
}

func (f *Done) Synopsis() string {
	return "done Task"
}

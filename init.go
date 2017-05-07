package main

import (
  "os"
  "fmt"
)
type Init struct{}

func (f *Init) Help() string {
	return "app Init"
}

func (f *Init) Run(args []string) int {
	root := GetRoot()
	os.Mkdir(root, 0777)
	taskRoot := root + "/Tasks"
	os.Mkdir(taskRoot, 0777)
	os.Mkdir(taskRoot+"/Finished", 0777)
	os.Mkdir(taskRoot+"/Have", 0777)
	fmt.Println("Make Tasks!")
	fmt.Println("Let's start to manage tasks on prompt")
	return 0
}

func (f *Init) Synopsis() string {
	return "Start Management Task"
}

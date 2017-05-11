package main

import (
	"fmt"
	"os"
)

type Init struct{}

func (f *Init) Help() string {
	return "app Init"
}

func (f *Init) Run(args []string) int {
	os.Mkdir(root.root, 0777)
	os.Mkdir(root.troot, 0777)
	os.Mkdir(root.finished, 0777)
	os.Mkdir(root.have, 0777)
	fmt.Println("Make Tasks!")
	fmt.Println("Let's start to manage tasks on prompt")
	return 0
}

func (f *Init) Synopsis() string {
	return "Start Management Task"
}

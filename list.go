package main

import (
	"fmt"
	"strings"
)

type List struct{}

func (f *List) Help() string {
	return "app List"
}

func (f *List) Run(args []string) int {
	root := GetRoot()
	roop := root.GetList(1)
	for _, v := range roop {
		taskname := strings.Split(v, ".json")
		fmt.Println(taskname[0])
	}
	return 0
}

func (f *List) Synopsis() string {
	return "display of all task name"
}

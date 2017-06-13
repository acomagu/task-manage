package main

import (
	"os"
)

type Show struct{}

func (f *Show) Help() string {
	return "app Show"
}

func (f *Show) Run(args []string) int {
	if len(os.Args) < 3 {
		roop := root.GetListHave()
		for _, v := range roop {
			TaskPrint(root.have + v)
		}
	} else {
		TaskPrint(filepath.Json(root.have, args[0]+".json"))
	}
	return 0
}

func (f *Show) Synopsis() string {
	return "display of task or all task"
}

package main

import (
	"fmt"
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
		TaskPrint(root.have + args[0] + ".json")
	}
	return 0
}

func (f *Show) Synopsis() string {
	return "display of task or all task"
}

func TaskPrint(task string) {
	data := FindTask(task)
	fmt.Println("Title : ", data.Title)
	fmt.Println("Content : ", data.Content)
	fmt.Println("Dead Line : ", data.DeadLine.Format("2006-01-02"))
	if data.DoneTime != data.DeadLine {
		fmt.Println("Done Time : ", data.DoneTime.Format("2006-01-02"))
	}
	fmt.Println("---------------------------------------------------------------")
}

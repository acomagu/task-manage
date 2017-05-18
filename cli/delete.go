package main

import (
	"fmt"
	"os"
)

type Delete struct{}

func (f *Delete) Help() string {
	return "task-manage delete <-h have, -f finished>"
}

func (f *Delete) Run(args []string) int {
	deleteRootPath := GetDeleteRootPath(args[1])
	if err := os.Remove(deleteRootPath + args[0] + ".json"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Delete ", args[0])
	}
	return 0
}

func (f *Delete) Synopsis() string {
	return "display of all tasks"
}

func GetDeleteRootPath(s string) string {
	if s == "-o" {
		return root.have
	} else {
		return root.finished
	}
}

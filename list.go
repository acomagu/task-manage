package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type List struct{}

func (f *List) Help() string {
	return "app List"
}

func (f *List) Run(args []string) int {
	root := GetRoot() + "/Tasks/"
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			rel, err := filepath.Rel(root, path)
			taskname := strings.Split(rel, "_Task")
			fmt.Println(taskname[0])
			return nil
		})
	if err != nil {
		fmt.Println(1, err)
	}
	return 0
}

func (f *List) Synopsis() string {
	return "display of all task name"
}

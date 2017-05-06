package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Show struct{}

func (f *Show) Help() string {
	return "app Show"
}

func (f *Show) Run(args []string) int {
	path := os.Getenv("GOPATH")
	root := path + "/src/github.com/task-manager/Tasks/Have/"
	if len(os.Args) < 3 {
		err := filepath.Walk(root,
			func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				rel, err := filepath.Rel(root, path)
				printj(root + rel)
				return nil
			})
		if err != nil {
			fmt.Println(1, err)
		}
	} else {
		printj(root + os.Args[2] + "_Task.json")
	}
	return 0
}

func (f *Show) Synopsis() string {
	return "Print \"Show!\""
}

func printj(root string) {
	bytes, err := ioutil.ReadFile(root)
	if err != nil {
		log.Fatal(err)
	}
	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Title : ", data.Title)
	fmt.Println("Content : ", data.Content)
	fmt.Println("Dead Line : ", data.DeadLine.Format("2006-01-02"))
	fmt.Println("---------------------------------------------------------------")
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Show struct{}

func (f *Show) Help() string {
	return "app Show"
}

func (f *Show) Run(args []string) int {
	path := os.Getenv("GOPATH")
	root := path + "/src/github.com/task-manager/Tasks/Have/"
	bytes, err := ioutil.ReadFile(root + os.Args[2] + "_Task.json")
	if err != nil {
		log.Fatal(err)
	}
	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Title : ", data.Title)
	fmt.Println("Content : ", data.Content)
	fmt.Println("Dead Line : ", data.DeadLine.Format("2017-01-01"))
	return 0
}

func (f *Show) Synopsis() string {
	return "Print \"Show!\""
}

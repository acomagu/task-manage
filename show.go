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
	root := GetRoot()
	if len(os.Args) < 3 {
		roop := root.GetList(2)
		for _, v := range roop {
			printj(root.have + v)
		}
	} else {
		printj(root.have + os.Args[2] + ".json")
	}
	return 0
}

func (f *Show) Synopsis() string {
	return "display of task or all task"
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

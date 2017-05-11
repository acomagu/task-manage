package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Edit struct{}

func (f *Edit) Help() string {
	return "task-manage edit <task Title> <task New Title> <task New Content> <task Nwe days>"
}

func (f *Edit) Run(args []string) int {
	root := NewRoot()
	task := root.have + os.Args[2] + ".json"
	bytes, err := ioutil.ReadFile(task)
	if err != nil {
		log.Fatal(err)
	}
	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	if err := os.Remove(task); err != nil {
		fmt.Println(err)
	}
	data.Title = os.Args[3]
	data.Content = os.Args[4]
	n, _ := strconv.Atoi(os.Args[5])
	end, _ := NewPoint(n)
	data.DeadLine = end
	fout, err := os.Create(root.have + os.Args[3] + ".json")
	if err != nil {
		fmt.Println(task, err)
	}
	outputJson, err := json.Marshal(&data)
	fout.Write([]byte(outputJson))
	if err != nil {
		panic(err)
	}
	defer fout.Close()
	Printj(root.have + os.Args[3] + ".json")
	return 0
}

func (f *Edit) Synopsis() string {
	return "edit task"
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Data struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	DeadLine  time.Time `json:"dead_line"`
}

type Add struct{}

func (f *Add) Help() string {
	return "You can add a task \n ./app add <title> <content> <Days>"
}

func (f *Add) Run(args []string) int {
	now := time.Now()
	i, _ := strconv.Atoi(os.Args[4])
	end := now.AddDate(0, 0, i)
	data := Data{
		os.Args[2],
		os.Args[3],
		now,
		end,
	}
	path := os.Getenv("GOPATH")
	root := path + "/src/github.com/task-manager/Tasks/Have/"
	userFile := root + os.Args[2] + "_Task.json"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
	}
	outputJson, err := json.Marshal(&data)
	if err != nil {
		panic(err)
	}
	defer fout.Close()
	fout.Write([]byte(outputJson))
	log.Println(userFile)
	return 0
}

func (f *Add) Synopsis() string {
	return "add Task"
}

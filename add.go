package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Data struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	DeadLine  time.Time `json:"deadline"`
}

type Add struct{}

func (f *Add) Help() string {
	return "You can add a task \n ./app add <title> <content> <Days>"
}

func (f *Add) Run(args []string) int {
	n, _ := strconv.Atoi(os.Args[4])
	end, now := NewPoint(n)
	data := Data{
		os.Args[2],
		os.Args[3],
		now,
		end,
	}
	root := NewRoot()
	task := root.have + os.Args[2] + ".json"
	fout, err := os.Create(task)
	if err != nil {
		fmt.Println(task, err)
	}
	outputJson, err := json.Marshal(&data)
	fout.Write([]byte(outputJson))
	if err != nil {
		panic(err)
	}
	defer fout.Close()
	Printj(task)
	return 0
}

func (f *Add) Synopsis() string {
	return "add Task"
}

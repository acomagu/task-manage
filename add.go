package main

import (
	"log"
	"strconv"
	"time"
)

type Data struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	DeadLine  time.Time `json:"deadline"`
	DoneTime  time.Time `json:"donetime"`
}

type Add struct{}

func (f *Add) Help() string {
	return "You can add a task \n ./app add <title> <content> <Days>"
}

func (f *Add) Run(args []string) int {
	n, err := strconv.Atoi(args[2])
	if err != nil {
		log.Println(err)
		return 1
	}
	end, now := NewPoint(n)
	data := Data{
		args[0],
		args[1],
		now,
		end,
		end,
	}
	creatore.Task(data, root.have)
	return 0
}

func (f *Add) Synopsis() string {
	return "add Task"
}

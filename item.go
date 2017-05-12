package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func NewRoot() Root {
	path := os.Getenv("HOME")
	r := path + "/.task-manage"
	root := Root{
		r,
		r + "/Tasks/",
		r + "/Tasks/Have/",
		r + "/Tasks/Finished/",
	}
	return root
}

func NewPoint(n int) (time.Time, time.Time) {
	now := time.Now()
	end := now.AddDate(0, 0, n)
	return end, now
}

func FindTask(task string) Data {
	bytes, err := ioutil.ReadFile(task)
	if err != nil {
		log.Fatal(err)
	}
	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	return data
}

type CreateFile struct{}

var creatore CreateFile

func (s CreateFile) Task(data Data, path string) {
	fout, err := os.Create(path + data.Title + ".json")
	if err != nil {
		fmt.Println(data.Title, err)
	}
	outputJson, err := json.Marshal(&data)
	fout.Write([]byte(outputJson))
	if err != nil {
		panic(err)
	}
	defer fout.Close()
	Printj(path + data.Title + ".json")
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

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

func (c CreateFile) Task(data Data, path string) {
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
	TaskPrint(path + data.Title + ".json")
}

func (c CreateFile) Recode(recorde []Recode) {
	fout, err := os.Create(root.root + "/recode.json")
	if err != nil {
		fmt.Println("Recode: ", err)
	}
	outputJson, err := json.Marshal(&recorde)
	fout.Write([]byte(outputJson))
	if err != nil {
		panic(err)
	}
	defer fout.Close()
	fmt.Println("Recoding file")
}

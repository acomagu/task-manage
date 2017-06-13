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

func TaskPrint(task string) {
	data := FindTask(task)
	fmt.Println("Title : ", data.Title)
	fmt.Println("Content : ", data.Content)
	fmt.Println("Dead Line : ", data.DeadLine.Format("2006-01-02"))
	if data.DoneTime != data.DeadLine {
		fmt.Println("Done Time : ", data.DoneTime.Format("2006-01-02"))
	}
	fmt.Println("---------------------------------------------------------------")
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

func FindTaskin5days(tasks []string) [][]string {
	day := time.Now()
	var days5 [][]string

	for _, v := range tasks {
		data := FindTask(v)
		sub := data.DeadLine.Sub(day)
		subdays := int(sub.Hours())
		if 0 <= subdays/24 && subdays/24 < 5 {
			taskdays := []string{"", "", "", "", ""}
			taskdays[subdays/24] = data.Title
			days5 = append(days5, taskdays)
		}
	}
	return days5
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

func ArrayContains(arr []string, str string) (int, bool) {
	for i, v := range arr {
		if v == str {
			return i, true
		}
	}
	return -1, false
}


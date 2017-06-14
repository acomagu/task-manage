package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

func FindTask(task string) (Data, error) {
	f, err := os.Open(task)
	if err != nil {
		return Data{}, err
	}
	defer f.Close()

	var data Data
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&data); err != nil {
		return Data{}, err
	}
	return data, nil
}

func FindTaskin5days(tasks []string) [][]string {
	day := time.Now()
	var days5 [][]string

	for _, v := range tasks {
		data := FindTask(v)
		sub := data.DeadLine.Sub(day)
		subdays := int(sub.Hours())
		if 0 <= subdays/24 && subdays/24 < 5 {
			taskdays := make([]string, 5)
			taskdays[subdays/24] = data.Title
			days5 = append(days5, taskdays)
		}
	}
	return days5
}

type CreateFile struct{}

var creatore CreateFile

func (c CreateFile) Task(data Data, path string) error {
	fout, err := os.Create(filepath.Join(path, data.Title+".json"))
	if err != nil {
		return err
	}
	defer fout.Close()

	encoder := json.NewEncoder(fout)
	if err := encoder.Encode(&data); err != nil {
		return err
	}
	TaskPrint(filepath.Join(path, data.Title+".json"))

	return nil
}

func ArrayContains(arr []string, str string) (int, bool) {
	for i, v := range arr {
		if v == str {
			return i, true
		}
	}
	return -1, false
}

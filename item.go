package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Replay struct {
	Recode []Recode `json:"recode"`
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
func NewIp() Ip {
	bytes, err := ioutil.ReadFile(root.root + "/ip.json")
	if err != nil {
		log.Fatal(err)
	}
	var data Ip
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	return data
}

func NewRecode() Replay {
	bytes, err := ioutil.ReadFile(root.root + "/recode.json")
	if err != nil {
		log.Fatal(err)
	}
	var data []Recode
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	return Replay{data}
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

func (c CreateFile) IpMemory(ip Ip) {
	fout, err := os.Create(root.root + "/ip.json")
	if err != nil {
		fmt.Println("Ip: ", err)
	}
	outputJson, err := json.Marshal(&ip)
	fout.Write([]byte(outputJson))
	if err != nil {
		panic(err)
	}
	defer fout.Close()
	fmt.Println("Memory ip")
}

func ArrayContains(arr []string, str string) (int, bool) {
	for i, v := range arr {
		if v == str {
			return i, true
		}
	}
	return -1, false
}

func HttpPost(url string, recode Replay) error {
	outputJson, err := json.Marshal(&recode)
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(outputJson)),
	)
	if err != nil {
		return err
	}
	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Input struct {
	In string
}

type Output struct {
	Out string
}

type Data struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	DeadLine  time.Time `json:"deadline"`
	DoneTime  time.Time `json:"donetime"`
}

type Recode struct {
	Done bool `json:"done"`
	Data Data `json:"data"`
}

func jsonHandleFunc(rw http.ResponseWriter, req *http.Request) {
	output := []Recode{}
	defer func() {
		outjson, e := json.Marshal(output)
		if e != nil {
			fmt.Println(e)
		}
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(outjson))
	}()

	if req.Method == "GET" {
		data := FindTask("./recode.json")
		outjson, e := json.Marshal(data)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Fprint(rw, string(outjson))
		return
	}
	body, e := ioutil.ReadAll(req.Body)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	input := []Recode{}
	e = json.Unmarshal(body, &input)
	output = input
	Recodes(input)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Printf("%#v\n", input)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/json", jsonHandleFunc)
	http.ListenAndServe(":8080", nil)
}

func Recodes(recorde []Recode) {
	fout, err := os.Create("recode.json")
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

func FindTask(task string) []Recode {
	bytes, err := ioutil.ReadFile(task)
	if err != nil {
		log.Fatal(err)
	}
	var data []Recode
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	return data
}

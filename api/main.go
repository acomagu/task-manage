package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"time"
)

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

type Replay struct {
	Recode []Recode `json:"recode"`
}

func jsonHandleFunc(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", req.RemoteAddr)
	rw.Header().Set("Access-Control-Allow-Credentials", "true")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	rw.Header().Set("Content-Type", "application/json")
	var output Replay
	// defer func() {
	// 	outjson, e := json.Marshal(output)
	// 	if e != nil {
	// 		fmt.Println(e)
	// 	}
	// 	fmt.Print("TTTTTTTTTTTT")
	// 	fmt.Fprint(rw, string(outjson))
	// }()

	if req.Method == "GET" {
		data := FindTask("./recode.json")
		json.NewEncoder(rw).Encode(data)
		fmt.Print("SSSSSSSSSSSS")
		return
	}
	body, e := ioutil.ReadAll(req.Body)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	if req.Method == "POST" {
		input := Replay{}
		e = json.Unmarshal(body, &input)
		output = input
		Recodes(input)
		if e != nil {
			fmt.Println(e.Error())
			return
		}
		json.NewEncoder(rw).Encode(output)
		fmt.Printf("%#v\n", input)
		return
	}
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}
	http.HandleFunc("/", jsonHandleFunc)
	fcgi.Serve(l, nil)
}

func Recodes(recorde Replay) {
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

func FindTask(task string) Replay {
	bytes, err := ioutil.ReadFile(task)
	if err != nil {
		log.Fatal(err)
	}
	var data Replay
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	return data
}

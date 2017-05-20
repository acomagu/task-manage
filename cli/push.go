package main

import "fmt"

type Ip struct {
	Ip string `json:"ip"`
}
type Push struct{}

func (f *Push) Help() string {
	return "app Push"
}

func (f *Push) Run(args []string) int {
	if num, g := ArrayContains(args, "-ip"); g {
		ip := Ip{args[num+1]}
		creatore.IpMemory(ip)
	}
	recode := NewRecode()
	ip := NewIp()
	err := HttpPost(ip.Ip, recode)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Push " + ip.Ip)
	}
	return 0
}

func (f *Push) Synopsis() string {
	return "push recode"
}

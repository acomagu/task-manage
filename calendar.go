package main

import (
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Calender struct{}

func (f *Calender) Help() string {
	return "app Calender"
}

func (f *Calender) Run(args []string) int {
	roop := root.GetListHave()
	var tasks []string
	for _, v := range roop {
		tasks = append(tasks, root.have+v)
	}
	day := time.Now()
	const layout = "02"
	today := string(day.Format(layout))

	daystrings := []string{}
	for i := 0; i < 5; i++ {
		day, _ := strconv.Atoi(today)
		day = day + i
		k := strconv.Itoa(day)
		daystrings = append(daystrings, k)
	}

	data := FindTaskin5days(tasks)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(daystrings)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	return 0
}

func (f *Calender) Synopsis() string {
	return "Start Management Task"
}

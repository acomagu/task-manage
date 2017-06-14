package main

import (
	"fmt"
)

func printTask(task Task) {
	fmt.Println("Title : ", task.Title)
	fmt.Println("Content : ", task.Content)
	fmt.Println("Dead Line : ", task.Deadline.Format("2006-01-02"))
	if task.FinishedAt != task.Deadline {
		fmt.Println("Done Time : ", task.FinishedAt.Format("2006-01-02"))
	}
	fmt.Println("---------------------------------------------------------------")
}

func ArrayContains(arr []string, str string) (int, bool) {
	for i, v := range arr {
		if v == str {
			return i, true
		}
	}
	return -1, false
}

package main

import "time"

type Task struct {
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	Deadline   time.Time `json:"deadline"`
	FinishedAt time.Time `json:"finished_at"`
}

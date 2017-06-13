package main

import (
	"os"
	"fmt"
	"path/filepath"
)

type Root struct {
	root     string
	troot    string
	have     string
	finished string
}

func NewRoot() Root {
	path := os.Getenv("HOME")
	r := path + "/.task-manage"
	root := Root{
		r,
		r + "/Tasks/",
		r + "/Tasks/Have/",
		r + "/Tasks/Finished/",
	}
	return root
}

var root = NewRoot()

func (r Root) GetRootTasks() string {
	return r.troot
}

func (r Root) GetListTasks() TaskList {
	return r.GetList(r.GetRootTasks())
}
func (r Root) GetRootFinished() string {
	return r.finished
}
func (r Root) GetListFinished() TaskList {
	return r.GetList(r.GetRootFinished())
}
func (r Root) GetRootHave() string {
	return r.have
}

func (r Root) GetListHave() TaskList {
	return r.GetList(r.GetRootHave())
}
func (r Root) GetList(rootpath string) TaskList {
	var result []string
	err := filepath.Walk(rootpath,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			rel, err := filepath.Rel(rootpath, path)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, rel)
			return nil
		})
	if err != nil {
		fmt.Println("Walk", err)
	}
	return result
}

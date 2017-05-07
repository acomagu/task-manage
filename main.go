package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

type Root struct {
	root     string
	troot    string
	have     string
	finished string
}

func (r Root) GetList(n int) []string {
	var result []string
	var rpath string
	switch n {
	case 0:
		rpath = r.root
	case 1:
		rpath = r.troot
	case 2:
		rpath = r.have
	case 3:
		rpath = r.finished
	default:
		rpath = r.root
	}
	err := filepath.Walk(rpath,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			rel, err := filepath.Rel(rpath, path)
			result = append(result, rel)
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func GetRoot() Root {
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

func main() {
	c := cli.NewCLI("app", "1.0.0")

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &Init{}, nil
		},
		"list": func() (cli.Command, error) {
			return &List{}, nil
		},
		"add": func() (cli.Command, error) {
			return &Add{}, nil
		},
		"show": func() (cli.Command, error) {
			return &Show{}, nil
		},
		"done": func() (cli.Command, error) {
			return &Done{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}

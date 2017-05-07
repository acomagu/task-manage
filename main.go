package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func GetRoot() string {
	path := os.Getenv("HOME")
	root := path + "/.task-manage"
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

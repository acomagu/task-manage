package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

var db DB

func main() {
	var err error
	db, err = newDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	c := createCLI()
	c.Args = os.Args[1:]

	initDataDir()
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}

func createCLI() *cli.CLI {
	c := cli.NewCLI("app", "1.0.0")

	c.Commands = map[string]cli.CommandFactory{
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
		"edit": func() (cli.Command, error) {
			return &Edit{}, nil
		},
		"delete": func() (cli.Command, error) {
			return &Delete{}, nil
		},
	}
	return c
}

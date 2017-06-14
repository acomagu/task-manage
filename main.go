package main

import (
	"fmt"
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

	ui := &cli.ColoredUi{
		InfoColor:  cli.UiColorBlue,
		ErrorColor: cli.UiColorRed,
		Ui: &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
	}

	c.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &List{ui}, nil
		},
		"add": func() (cli.Command, error) {
			return &Add{ui}, nil
		},
		"show": func() (cli.Command, error) {
			return &Show{ui}, nil
		},
		"finish": func() (cli.Command, error) {
			return &Done{ui}, nil
		},
		"edit": func() (cli.Command, error) {
			return &Edit{ui}, nil
		},
		"delete": func() (cli.Command, error) {
			return &Delete{ui}, nil
		},
	}
	return c
}

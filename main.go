package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

type Init struct{}

func (f *Init) Help() string {
	return "app Init"
}

func (f *Init) Run(args []string) int {
	path := os.Getenv("GOPATH")
	root := path + "/src/github.com/Yamashou/task-manage"
	root = root + "/Tasks"
	os.Mkdir(root, 0777)
	os.Mkdir(root+"/Finished", 0777)
	os.Mkdir(root+"/Have", 0777)
	log.Println("Make Tasks!")
	log.Println("Let's start to manage tasks on prompt")
	return 0
}

func (f *Init) Synopsis() string {
	return "Print \"Init!\""
}

func main() {
	// コマンドの名前とバージョンを指定
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
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}

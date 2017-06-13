package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

type Delete struct {
	ui cli.Ui
}

func (f *Delete) Help() string {
	return "task-manage delete <-h have, -f finished>"
}

func (f *Delete) Run(args []string) int {
	deleteRootPath := GetDeleteRootPath(args[1])
	if err := os.Remove(filepath.Join(deleteRootPath, args[0]+".json")); err != nil {
		f.ui.Error(fmt.Sprint(err))
		return 1
	}
	f.ui.Info(fmt.Sprintf("Deleted: %s", args[0]))
	return 0
}

func (f *Delete) Synopsis() string {
	return "display of all tasks"
}

func GetDeleteRootPath(s string) string {
	if s == "-o" {
		return root.have
	}
	return root.finished
}

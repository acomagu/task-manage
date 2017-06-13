package main

import (
	"fmt"
	"os"
)

func initDataDir(args []string) int {
	os.Mkdir(root.root, 0777)
	os.Mkdir(root.troot, 0777)
	os.Mkdir(root.finished, 0777)
	os.Mkdir(root.have, 0777)
	fmt.Println("Make Tasks!")
	fmt.Println("Let's start to manage tasks on prompt")
	return 0
}

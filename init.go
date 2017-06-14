package main

import (
	"os"
)

func initDataDir() {
	os.Mkdir(db.path.Root(), 0777)
	os.Mkdir(db.path.Ongoing(), 0777)
	os.Mkdir(db.path.Finished(), 0777)
}

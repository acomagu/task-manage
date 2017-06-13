package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type DB struct {
	path DBPath
}

func newDB() (DB, error) {
	dbPath, err := newDBPath()
	if err != nil {
		return DB{}, err
	}

	return DB{
		path: dbPath,
	}, nil
}

func (db DB) All() TaskList {
	return db.collect(db.path.Root())
}

func (db DB) Finished() TaskList {
	return db.collect(db.path.Finished())
}

func (db DB) Ongoing() TaskList {
	return db.collect(db.path.Ongoing())
}

// collect lists all file paths under the rootpath.
func (db DB) collect(rootpath string) TaskList {
	var result TaskList
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(rootpath, path)
		if err != nil {
			return err
		}

		result = append(result, rel)
		return nil
	})

	if err != nil {
		fmt.Println("Walk", err)
	}
	return result
}

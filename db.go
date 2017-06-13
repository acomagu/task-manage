package main

import (
	"os"
	"fmt"
	"path/filepath"
	"crypto/sha512"
	"encoding/json"
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

func (db DB) Store(task Task) error {
	id := fmt.Sprintf("%x", sha512.Sum512([]byte(task.Title)))[:10]
	filename := fmt.Sprintf("%s.json", id)
	fout, err := os.Create(filepath.Join(db.path.Ongoing(), filename))
	if err != nil {
		return err
	}

	outputJson, err := json.Marshal(&task)
	fout.Write([]byte(outputJson))
	defer fout.Close()
	if err != nil {
		return err
	}

	// TaskPrint(path + data.Title + ".json")
	return nil
}

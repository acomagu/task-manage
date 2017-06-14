package main

import (
	"fmt"
	"os"
	"path/filepath"
	"crypto/sha512"
	"encoding/json"
	"time"
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

func (db DB) readFrom(path string) (Task, error) {
	f, err := os.Open(path)
	if err != nil {
		return Task{}, err
	}
	defer f.Close()

	var task Task
	if err := json.NewDecoder(f).Decode(&task); err != nil {
		return Task{}, err
	}
	return task, nil
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

		result = append(result, filepath.Join(db.path.Root(), rel))
		return nil
	})

	if err != nil {
		fmt.Println("Walk", err)
	}
	return result
}

func (db DB) Store(task Task) error {
	db.createOf(task, ongoing)

	// TaskPrint(path + data.Title + ".json")
	return nil
}

func (db DB) Finish(title string) error {
	task, err := db.readFrom(db.calcFilePath(title, ongoing))
	if err != nil {
		return err
	}

	err = db.deleteOf(title, ongoing)
	if err != nil {
		return err
	}

	task.FinishedAt = time.Now()

	return db.createOf(task, finished)
}

func (db DB) deleteOf(title string, state TaskState) error {
	path := db.calcFilePath(title, state)
	return os.Remove(path)
}

func (db DB) calcFilePath(title string, state TaskState) string {
	filename := db.calcFileName(title)
	return filepath.Join(db.stateDirPath(state), filename)
}

func (db DB) calcFileName(title string) string {
	id := fmt.Sprintf("%x", sha512.Sum512([]byte(title)))[:10]
	return fmt.Sprintf("%s.json", id)
}

func (db DB) createOf(task Task, state TaskState) error {
	fout, err := os.Create(db.calcFilePath(task.Title, ongoing))
	if err != nil {
		return err
	}
	defer fout.Close()

	return json.NewEncoder(fout).Encode(&task)
	// TaskPrint(filepath.Join(path, data.Title+".json"))
}

func (db DB) stateDirPath(state TaskState) string {
	if state == ongoing {
		return db.path.Ongoing()
	} else if state == finished {
		return db.path.Finished()
	}
	return ""
}

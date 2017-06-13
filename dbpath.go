package main

import (
	"path/filepath"
	"os"

	"github.com/mitchellh/go-homedir"
)

type DBPath struct {
	root     string
}

func newDBPath() (DBPath, error) {
	dataHome, err := getDataHome()
	if err != nil {
		return DBPath{}, err
	}

	return DBPath{
		root: filepath.Join(dataHome, "task-manage"),
	}, nil
}

func getDataHome() (string, error) {
	xdgDataHome := os.Getenv("XDG_DATA_HOME")

	if xdgDataHome != "" {
		return filepath.Clean(xdgDataHome), nil
	}

	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".local", "share"), nil
}

func (dbPath DBPath) Root() string {
	return dbPath.root
}

func (dbPath DBPath) Finished() string {
	return filepath.Join(dbPath.root, "finished")
}

func (dbPath DBPath) Ongoing() string {
	return filepath.Join(dbPath.root, "ongoing")
}

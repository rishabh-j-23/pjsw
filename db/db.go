package db

import (
	"database/sql"
	"io/fs"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

const dbFileName = "project_database.db"

func CreateDatabase() (*Projects, error) {
	homeDir, err := os.UserHomeDir()
	projectDir := path.Join(homeDir, ".pjsw")

	_, err = os.Stat(projectDir)
	if err != nil {
		err = os.Mkdir(projectDir, fs.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	dbFile := path.Join(projectDir, dbFileName)

	// give db file read and write permissions
	os.Chmod(dbFile, 0666)

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}
	return &Projects{db}, err
}

func InitDatabase() *Projects {
	project, err := CreateDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	project.CreateTable()

	return project
}

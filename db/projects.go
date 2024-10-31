package db

import (
	"database/sql"
	"fmt"
)

const projectTable string = "projects"

type Projects struct {
	db *sql.DB
}

func (project *Projects) CreateTable() (sql.Result, error) {
	const createDb string = `
        CREATE TABLE IF NOT EXISTS projects (
            projectName VARCHAR(20) PRIMARY KEY,
            projectPath TEXT
        );`

	result, err := project.db.Exec(createDb)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (project *Projects) InsertIntoTable(projectName string, projectPath string) (sql.Result, error) {
	const insert string = `INSERT INTO projects VALUES (?, ?);`

	result, err := project.db.Exec(insert, projectName, projectPath)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (project *Projects) GetPathByProject(name string) (string, error) {
	var path string
	const getPath = `SELECT projectPath FROM projects WHERE projectName=?`
	err := project.db.QueryRow(getPath, name).Scan(&path)
	if err != nil {
		return "", fmt.Errorf("error: Project '%s' does not exist", name)
	}

	return path, nil
}

func (p *Projects) GetAll() (map[string]string, error) {
	query := "SELECT * FROM " + projectTable + ";"

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	projectsTable := make(map[string]string)

	for rows.Next() {
		var name, path string

		if err := rows.Scan(&name, &path); err != nil {
			return nil, err
		}

		projectsTable[name] = path
	}

	return projectsTable, nil
}

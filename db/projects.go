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

func (project *Projects) DeleteProjectByName(name string) (string, error) {
	var existingName string
	row := project.db.QueryRow("SELECT projectName FROM projects WHERE projectName=?", name)
	err := row.Scan(&existingName)
	if err != nil {
		// If no rows are found, `Scan` will return `sql.ErrNoRows`
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("error: Project '%s' does not exist", name)
		}
		return "", err // Return any other error that might occur
	}

	const deletePath = `DELETE FROM projects WHERE projectName=?`
	_, err = project.db.Exec(deletePath, name)
	if err != nil {
		return "", err
	}

	return name, nil
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

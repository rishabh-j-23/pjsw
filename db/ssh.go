package db

import (
	"database/sql"
	"fmt"
)

type SSHConnection struct {
	Name         string
	User         string
	Host         string
	Port         string
	IdentityFile string
}

func (p *Projects) CreateSSHTable() (sql.Result, error) {
	const createDb string = `
        CREATE TABLE IF NOT EXISTS ssh_connections (
            name VARCHAR(50) PRIMARY KEY,
            user VARCHAR(50),
            host VARCHAR(100),
            port VARCHAR(10),
            identity_file TEXT
        );`

	result, err := p.db.Exec(createDb)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Projects) AddSSHConnection(conn SSHConnection) (sql.Result, error) {
	const insert string = `INSERT INTO ssh_connections (name, user, host, port, identity_file) VALUES (?, ?, ?, ?, ?);`

	result, err := p.db.Exec(insert, conn.Name, conn.User, conn.Host, conn.Port, conn.IdentityFile)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Projects) GetSSHConnection(name string) (*SSHConnection, error) {
	var conn SSHConnection
	const getConn = `SELECT name, user, host, port, identity_file FROM ssh_connections WHERE name=?`
	
	err := p.db.QueryRow(getConn, name).Scan(&conn.Name, &conn.User, &conn.Host, &conn.Port, &conn.IdentityFile)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error: SSH connection '%s' does not exist", name)
		}
		return nil, err
	}

	return &conn, nil
}

func (p *Projects) DeleteSSHConnection(name string) (string, error) {
	var existingName string
	row := p.db.QueryRow("SELECT name FROM ssh_connections WHERE name=?", name)
	err := row.Scan(&existingName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("error: SSH connection '%s' does not exist", name)
		}
		return "", err
	}

	const deleteConn = `DELETE FROM ssh_connections WHERE name=?`
	_, err = p.db.Exec(deleteConn, name)
	if err != nil {
		return "", err
	}

	return name, nil
}

func (p *Projects) GetAllSSHConnections() ([]SSHConnection, error) {
	query := `SELECT name, user, host, port, identity_file FROM ssh_connections`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []SSHConnection
	for rows.Next() {
		var conn SSHConnection
		if err := rows.Scan(&conn.Name, &conn.User, &conn.Host, &conn.Port, &conn.IdentityFile); err != nil {
			return nil, err
		}
		connections = append(connections, conn)
	}

	return connections, nil
}

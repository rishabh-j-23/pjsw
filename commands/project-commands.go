package commands

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"pjsw/db"
	"strings"
	"text/tabwriter"
)

func switchProject(projectName string, db *db.Projects) {
	path, err := db.GetPathByProject(projectName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(path)
}

func addProject(name string, path string, db *db.Projects) {
	if path == "." {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalln("error getting working directory", err)
		}
		path = wd
	}

	_, err := db.InsertIntoTable(name, path)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetAllProjects(db *db.Projects) {
	data, err := db.GetAll()
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize a new tabwriter
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	defer writer.Flush()

	// Print headers
	fmt.Fprintln(writer, "project-name\t   path")

	// Print rows
	for k, v := range data {
		fmt.Fprintf(writer, "%s\t   %s\n", k, v)
	}
}

func DeleteProjectByName(name string, db *db.Projects) {
	delPath, err := db.DeleteProjectByName(name)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("'%s' removed successfully", delPath)
}

// FzfSelect retrieves all project keys from the database, pipes them to fzf for selection,
// and returns the selected project key.
func FzfSelectProject(db *db.Projects) (string, error) {
	// Retrieve all project data
	data, err := db.GetAll()
	if err != nil {
		return "", err
	}

	// Format data as newline-separated keys (assuming data is a slice of structs with a Key field)
	var keyBuffer bytes.Buffer
	for name := range data {
		// Assuming project has a Key field; adjust based on actual struct
		keyBuffer.WriteString(name + "\n")
	}

	// Set up fzf command
	cmd := exec.Command("fzf", "--no-sort", "--reverse", "--height=50%")
	cmd.Stdin = bytes.NewReader(keyBuffer.Bytes())
	cmd.Stderr = os.Stderr // Forward fzf errors to terminal

	// Run fzf and capture output
	output, err := cmd.Output()
	if err != nil {
		// Handle case where user exits fzf without selecting (e.g., Ctrl+C)
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 130 {
			return "", fmt.Errorf("no project selected")
		}
		return "", fmt.Errorf("fzf error: %w", err)
	}

	// Parse and trim the selected project key
	selectedProject := strings.TrimSpace(string(output))
	if selectedProject == "" {
		return "", fmt.Errorf("no project selected")
	}

	return data[selectedProject], nil
}

package commands

import (
	"fmt"
	"log"
	"os"
	"pjsw/db"
	"text/tabwriter"

	"golang.design/x/clipboard"
)

func switchProject(projectName string, db *db.Projects) {
	path, err := db.GetPathByProject(projectName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cdCommand := fmt.Sprintf("cd %s", path)
	clipboard.Write(clipboard.FmtText, []byte(cdCommand))
	fmt.Printf("'%s' copied to clipboard.\n", cdCommand)
	fmt.Println("Paste the command to switch directory")
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

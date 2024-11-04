package commands

import (
	"fmt"
	"os"
	"pjsw/db"
)

func Exec(args []string, db *db.Projects) {
	switch args[1] {
	case "sw":
		switchProject(args[2], db)
	case "add":
		if len(args) < 4 {
			fmt.Println("pjsw add <name> <path>")
			os.Exit(1)
		}
		addProject(args[2], args[3], db)
	case "getall", "all":
		GetAllProjects(db)
	case "help":
		PrintHelp()
	case "rm", "remove", "delete":
		DeleteProjectByName(args[2], db)
	default:
		fmt.Println("Command does not exists. use 'pjsw help' for more detail")
	}
}

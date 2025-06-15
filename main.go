package main

import (
	"log"
	"pjsw/commands"
	"pjsw/db"
	"pjsw/util"
)

func main() {
	projects := db.InitDatabase()

	args, err := util.GetArgs()
	if err != nil {
		log.Fatalln(err)
	}

	commands.Exec(args, projects)
}

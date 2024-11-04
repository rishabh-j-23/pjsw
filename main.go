package main

import (
	"log"
	"pjsw/commands"
	"pjsw/db"
	"pjsw/util"
)

func main() {
	db.InitDatabase()

	args, err := util.GetArgs()
	if err != nil {
		log.Fatalln(err)
	}
	db, err := db.CreateDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	commands.Exec(args, db)
}

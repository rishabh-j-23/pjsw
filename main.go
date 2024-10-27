package main

import (
	"flag"
	"fmt"
	"os"
	"pjsw/util"
)

func main() {
	arg, err := util.GetArg()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(arg)

	flag.String("greeting", "Hello", "Custom greeting message")

	// Parse flags; flag package only parses after the main command
	flag.CommandLine.Parse(os.Args[2:])

	fmt.Println(flag.Args())
}

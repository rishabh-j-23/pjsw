package commands

import "fmt"

func Cmdformat(mainCmd string, subCmd string, additionalParams []string) string {
	additionalParamsStr := ""
	if len(additionalParams) > 0 {
		for _, param := range additionalParams {
			additionalParamsStr += fmt.Sprintf("<%s> ", param)
		}
	}

	if subCmd != "" {
		subCmd = "<" + subCmd + ">"
	}

	return fmt.Sprintf("pjsw %s %s %s", mainCmd, subCmd, additionalParamsStr)
}

func PrintHelp() {
	CmdDetails([]string{"help"}, "Prints this message")
	CmdDetails([]string{"sw", "name"}, "switch(copies) dir of projects to clipboard. Command copied in clipboard is in format cd <projectpath>")
	CmdDetails([]string{"add", "name", "path"}, "adds project name and path to database. Use unique project name to avoid errors. Use '.' to refer to current working directory")
	CmdDetails([]string{"getall | all"}, "Prints all added projects name and path")
	CmdDetails([]string{"rm | remove | delete", "name"}, "Deletes path for added project. need to pass project 'name'")
}

func CmdDetails(cmds []string, cmdDescription string) {
	mainCmd := cmds[0]
	subCmd := ""
	addCmd := []string{}

	if len(cmds) > 1 {
		subCmd = cmds[1]
		addCmd = cmds[2:]
	}

	cmdFmt := Cmdformat(mainCmd, subCmd, addCmd)

	fmt.Printf(`
Command: %s 
Description: 
        %s 
`, cmdFmt, cmdDescription)
}

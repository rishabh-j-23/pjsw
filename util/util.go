package util

import (
	"errors"
	"os"
	"path/filepath"
)

func GetCurrentWorkingDir() (string, error) {
	return filepath.Abs(filepath.Dir("./"))
}

func GetArg() (string, error) {
	args := os.Args

	if len(args) == 1 {
		return "", errors.New("0 arguments passed. Need 1 argument. Get more detail with pjsw help <command>")
	}

	if len(args) > 2 && args[1] != "help" {
		return "", errors.New("Too many arguments passed (>1)")
	}

	return args[1], nil
}

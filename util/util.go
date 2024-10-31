package util

import (
	"errors"
	"os"
	"path/filepath"
)

func GetCurrentWorkingDir() (string, error) {
	return filepath.Abs(filepath.Dir("./"))
}

func GetArgs() ([]string, error) {
	args := os.Args

	if len(args) == 1 {
		return nil, errors.New("0 arguments passed. Need 1 argument. Get more detail with pjsw help")
	}

	return args, nil
}

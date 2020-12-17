package utils

import (
	"os"
)

func FileIsExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}

func CreateFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

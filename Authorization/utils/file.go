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

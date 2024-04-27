package util

import (
	"errors"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func RemoveFile(fileName string) error {
	if Exists(fileName) {
		if err := os.Remove(fileName); err != nil {
			return errors.New("failed to delete file")
		}
	}
	return nil
}

package util

import (
	"errors"
	"h-ui/model/constant"
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
			return errors.New(constant.RemoveFileError)
		}
	}
	return nil
}

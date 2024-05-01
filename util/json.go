package util

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"h-ui/model/constant"
	"os"
)

func ExportJson(filePath string, data any) error {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		logrus.Errorf("ExportJson create json file err filePath: %s err: %v", filePath, err)
		return errors.New(constant.SysError)
	}

	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		logrus.Errorf("ExportJson json Marshal err filePath: %s err: %v", filePath, err)
		return errors.New(constant.SysError)
	}
	_, err = file.Write(jsonBytes)
	if err != nil {
		logrus.Errorf("ExportJson writer WriteString err filePath: %s err: %v", filePath, err)
		return errors.New(constant.SysError)
	}
	return nil
}

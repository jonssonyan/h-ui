package util

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"h-ui/model/constant"
	"os"
)

// ExportFile t 0/json 1/yaml
func ExportFile(filePath string, data any, t int) error {
	file, err := os.Create(filePath)
	if err != nil {
		logrus.Errorf("ExportFile create file err filePath: %s err: %v", filePath, err)
		return errors.New(constant.SysError)
	}
	defer file.Close()
	var bytes []byte
	if t == 0 {
		bytes, err = json.MarshalIndent(data, "", "    ")
		if err != nil {
			logrus.Errorf("ExportFile Marshal json err filePath: %s err: %v", filePath, err)
			return errors.New(constant.SysError)
		}
	} else if t == 1 {
		bytes, err = yaml.Marshal(&data)
		if err != nil {
			logrus.Errorf("ExportFile Marshal yaml err filePath: %s err: %v", filePath, err)
			return errors.New(constant.SysError)
		}
	}
	_, err = file.Write(bytes)
	if err != nil {
		logrus.Errorf("ExportFile writer WriteString err filePath: %s err: %v", filePath, err)
		return errors.New(constant.SysError)
	}
	return nil
}

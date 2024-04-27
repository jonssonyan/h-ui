package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"os"
)

func Hysteria2Auth(pass string) (string, error) {
	account, err := dao.GetAccount("deleted = 0 and pass = ? and CURRENT_TIMESTAMP < expire_time and (quota < 0 or quota > download + upload) ", pass)
	if err != nil {
		return "", err
	}
	return *account.Username, nil
}

func SetHysteria2ConfigYAML() error {
	hysteria2Config, err := dao.GetConfig(constant.Hysteria2Config)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fmt.Sprintf("%s/%s", constant.BinDir, constant.Hysteria2ConfigPath), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("create hysteria2 server config file err: %v", err)
		return err
	}
	_, err = file.WriteString(*hysteria2Config.Value)
	if err != nil {
		logrus.Errorf("hysteria2 config.json file write err: %v", err)
		return err
	}
	return nil
}

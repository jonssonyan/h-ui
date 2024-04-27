package service

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/proxy"
	"h-ui/util"
	"os"
)

func Hysteria2Auth(pass string) (string, error) {
	account, err := dao.GetAccount("deleted = 0 and pass = ? and CURRENT_TIMESTAMP < expire_time and (quota < 0 or quota > download + upload) ", pass)
	if err != nil {
		return "", err
	}
	return *account.Username, nil
}

func GetHysteria2Config() (bo.Hysteria2ServerConfig, error) {
	var serverConfig bo.Hysteria2ServerConfig
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Enable)
	if err != nil {
		return serverConfig, err
	}
	if err := yaml.Unmarshal([]byte(*config.Value), &serverConfig); err != nil {
		return serverConfig, err
	}
	return serverConfig, nil
}

func SetHysteria2ConfigYAML() error {
	hysteria2Config, err := dao.GetConfig(constant.Hysteria2Config)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(constant.Hysteria2ConfigPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
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

func InitHysteria2() {
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Enable)
	if err != nil {
		logrus.Errorf("get hysteria2 enable config err: %v", err)
		return
	}

	if *config.Value == "1" {
		hysteria2Config, err := GetHysteria2Config()
		if err != nil {
			logrus.Errorf("get hysteria2 config err: %v", err)
			return
		}
		if err = SetHysteria2ConfigYAML(); err != nil {
			logrus.Errorf("set hysteria2 config.yaml err: %v", err)
			return
		}
		hysteria2Instance := proxy.NewHysteria2Instance(hysteria2Config.Listen, util.GetHysteria2BinPath(), constant.Hysteria2ConfigPath)
		if err = hysteria2Instance.StartHysteria2(); err != nil {
			logrus.Errorf("start hysteria2 err: %v", err)
			return
		}
	}
}

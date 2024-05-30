package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/proxy"
	"h-ui/util"
	"os"
)

func InitHysteria2() {
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Enable)
	if err != nil {
		logrus.Errorf("get hysteria2 enable config err: %v", err)
		return
	}

	if *config.Value == "1" {
		if err := StartHysteria2(); err != nil {
			return
		}
	}
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

func StartHysteria2() error {
	// 初始化 bin
	if !util.Exists(util.GetHysteria2BinPath()) {
		// 指定版本防止 api 不兼容
		if err := util.DownloadHysteria2("app/v2.4.4"); err != nil {
			return err
		}
	}
	// 初始化配置文件
	if err := SetHysteria2ConfigYAML(); err != nil {
		logrus.Errorf("set hysteria2 config.yaml err: %v", err)
		return errors.New("set hysteria2 config.yaml err")
	}
	hysteria2Instance := proxy.NewHysteria2Instance(util.GetHysteria2BinPath(), constant.Hysteria2ConfigPath)
	if err := hysteria2Instance.StartHysteria2(); err != nil {
		logrus.Errorf("start hysteria2 err: %v", err)
		return errors.New("start hysteria2 err")
	}
	return nil
}

func StopHysteria2() error {
	hysteria2Instance := proxy.NewHysteria2Instance("", constant.Hysteria2ConfigPath)
	if err := hysteria2Instance.StopHysteria2(); err != nil {
		logrus.Errorf("stop hysteria2 err: %v", err)
		return errors.New("stop hysteria2 err")
	}
	return nil
}

func RestartHysteria2() error {
	// 初始化 bin
	if !util.Exists(util.GetHysteria2BinPath()) {
		// 指定版本防止 api 不兼容
		if err := util.DownloadHysteria2("app/v2.4.4"); err != nil {
			return err
		}
	}
	// 初始化配置文件
	if err := SetHysteria2ConfigYAML(); err != nil {
		logrus.Errorf("set hysteria2 config.yaml err: %v", err)
		return errors.New("set hysteria2 config.yaml err")
	}
	hysteria2Instance := proxy.NewHysteria2Instance(util.GetHysteria2BinPath(), constant.Hysteria2ConfigPath)
	if err := hysteria2Instance.RestartHysteria2(); err != nil {
		logrus.Errorf("restart hysteria2 err: %v", err)
		return errors.New("restart hysteria2 err")
	}
	return nil
}

package service

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/proxy"
	"h-ui/util"
	"os"
	"strconv"
	"strings"
)

var hysteria2Instance *proxy.Hysteria2Process

func init() {
	hysteria2Instance = proxy.NewHysteria2Instance(util.GetHysteria2BinPath(), constant.Hysteria2ConfigPath)
}

func InitHysteria2() {
	// 初始化 bin
	if !util.Exists(util.GetHysteria2BinPath()) {
		// 下载最新版
		if err := util.DownloadHysteria2(""); err != nil {
			panic(fmt.Sprintf("download hysteris2 bin err: %v", err))
		}
	}

	config, err := dao.GetConfig("key = ?", constant.Hysteria2Enable)
	if err != nil {
		panic(fmt.Sprintf("get hysteria2 enable config err: %v", err))
	}

	if *config.Value == "1" {
		if err = StartHysteria2(); err != nil {
			panic(fmt.Sprintf("start hysteria2 server err: %v", err))
		}
	}
}

func SetHysteria2ConfigYAML() error {
	hysteria2Config, err := dao.GetConfig("key = ?", constant.Hysteria2Config)
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
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Config)
	if err != nil {
		return err
	}
	if config.Value == nil || *config.Value == "" {
		return errors.New("hysteria2 config is empty")
	}
	hysteria2Config, err := GetHysteria2Config()
	if err != nil {
		return err
	}

	// 代理端口
	listen, err := strconv.Atoi(strings.Trim(*hysteria2Config.Listen, ":"))
	if err != nil {
		return err
	}
	listenPortAvailable := util.IsPortAvailable(uint(listen), "udp")
	if !listenPortAvailable {
		return errors.New("hysteria port has been used")
	}

	// api 端口
	apiPort, err := strconv.Atoi(strings.Trim(*hysteria2Config.TrafficStats.Listen, ":"))
	if err != nil {
		return err
	}
	apiPortAvailable := util.IsPortAvailable(uint(apiPort), "tcp")
	if !apiPortAvailable {
		return errors.New("hysteria api port has been used")
	}

	// 初始化配置文件
	if err := SetHysteria2ConfigYAML(); err != nil {
		logrus.Errorf("set hysteria2 config.yaml err: %v", err)
		return errors.New("set hysteria2 config.yaml err")
	}
	if err := hysteria2Instance.StartHysteria2(); err != nil {
		logrus.Errorf("start hysteria2 err: %v", err)
		return errors.New("start hysteria2 err")
	}
	return nil
}

func StopHysteria2() error {
	if err := hysteria2Instance.StopHysteria2(); err != nil {
		logrus.Errorf("stop hysteria2 err: %v", err)
		return errors.New("stop hysteria2 err")
	}
	return nil
}

func RestartHysteria2() error {
	if err := StopHysteria2(); err != nil {
		logrus.Errorf("stop hysteria2 err: %v", err)
		return errors.New("stop hysteria2 err")
	}
	if err := StartHysteria2(); err != nil {
		logrus.Errorf("start hysteria2 err: %v", err)
		return errors.New("start hysteria2 err")
	}
	return nil
}

func Hysteria2Status() bool {
	return hysteria2Instance.IsRunning()
}

func ReleaseHysteria2() {
	if err := hysteria2Instance.Release(); err != nil {
		logrus.Errorf(err.Error())
	}
}

package service

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/model/vo"
	"h-ui/proxy"
	"h-ui/util"
	"os"
	"strconv"
	"strings"
)

func InitHysteria2() error {
	if !util.Exists(util.GetHysteria2BinPath()) {
		if err := util.DownloadHysteria2(""); err != nil {
			return errors.New(fmt.Sprintf("download hysteris2 bin err: %v", err))
		}
	}

	config, err := dao.GetConfig("key = ?", constant.Hysteria2Enable)
	if err != nil {
		return err
	}

	if *config.Value == "1" {
		if err = StartHysteria2(); err != nil {
			return err
		}
	}
	return nil
}

func SetHysteria2ConfigYAML() error {
	port, crtPath, keyPath, err := GetPortAndCert()
	if err != nil {
		return err
	}
	serverConfig, err := GetHysteria2Config()
	if err != nil {
		return err
	}
	if (crtPath != "" && keyPath != "") != (serverConfig.Auth != nil &&
		serverConfig.Auth.HTTP != nil &&
		serverConfig.Auth.HTTP.URL != nil &&
		strings.HasPrefix(*serverConfig.Auth.HTTP.URL, "https")) {
		protocol := "http"
		if crtPath != "" && keyPath != "" {
			protocol = "https"
		}
		url := fmt.Sprintf("%s://127.0.0.1:%d/hui/hysteria2/auth", protocol, port)
		serverConfig.Auth.HTTP.URL = &url
		if err := UpdateHysteria2Config(serverConfig); err != nil {
			return err
		}
	}

	hysteria2Config, err := yaml.Marshal(serverConfig)
	if err != nil {
		logrus.Errorf("marshal hysteria2 config err: %v", err)
		return err
	}

	file, err := os.OpenFile(constant.Hysteria2ConfigPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("create hysteria2 server config file err: %v", err)
		return err
	}
	_, err = file.WriteString(string(hysteria2Config))
	if err != nil {
		logrus.Errorf("hysteria2 config.json file write err: %v", err)
		return err
	}
	return nil
}

func Hysteria2IsRunning() bool {
	return proxy.NewHysteria2Instance().IsRunning()
}

func StartHysteria2() error {
	hysteria2Config, err := GetHysteria2Config()
	if err != nil {
		return err
	}

	if hysteria2Config.Listen == nil || *hysteria2Config.Listen == "" {
		return errors.New("hysteria2 config is empty")
	}

	listen, err := strconv.Atoi(strings.Trim(*hysteria2Config.Listen, ":"))
	if err != nil {
		return errors.New(fmt.Sprintf("hysteria port: %s is invalid", *hysteria2Config.Listen))
	}
	listenPortAvailable := util.IsPortAvailable(uint(listen), "udp")
	if !listenPortAvailable {
		return errors.New("hysteria port has been used")
	}

	apiPort, err := strconv.Atoi(strings.Trim(*hysteria2Config.TrafficStats.Listen, ":"))
	if err != nil {
		return errors.New(fmt.Sprintf("hysteria api port: %s is invalid", *hysteria2Config.Listen))
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
	if err := proxy.NewHysteria2Instance().StartHysteria2(); err != nil {
		logrus.Errorf("start hysteria2 err: %v", err)
		return errors.New("start hysteria2 err")
	}
	return nil
}

func StopHysteria2() error {
	if err := proxy.NewHysteria2Instance().StopHysteria2(); err != nil {
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

func ReleaseHysteria2() error {
	return proxy.NewHysteria2Instance().Release()
}

func Hysteria2AcmePath() (vo.Hysteria2AcmePathVo, error) {
	hysteria2AcmePathVo := vo.Hysteria2AcmePathVo{}
	hysteria2Config, err := GetHysteria2Config()
	if err != nil {
		return hysteria2AcmePathVo, err
	}
	if hysteria2Config.TLS != nil &&
		hysteria2Config.TLS.Cert != nil && *hysteria2Config.TLS.Cert != "" &&
		hysteria2Config.TLS.Key != nil && *hysteria2Config.TLS.Key != "" {
		if util.Exists(*hysteria2Config.TLS.Cert) && util.Exists(*hysteria2Config.TLS.Key) {
			hysteria2AcmePathVo.CrtPath = *hysteria2Config.TLS.Cert
			hysteria2AcmePathVo.KeyPath = *hysteria2Config.TLS.Key
			return hysteria2AcmePathVo, nil
		}
		return hysteria2AcmePathVo, errors.New("cert not found")
	} else if hysteria2Config.ACME != nil &&
		hysteria2Config.ACME.Domains != nil &&
		len(hysteria2Config.ACME.Domains) > 0 &&
		hysteria2Config.ACME.CA != nil &&
		*hysteria2Config.ACME.CA != "" &&
		hysteria2Config.ACME.Dir != nil &&
		*hysteria2Config.ACME.Dir != "" {
		acmeDir := *hysteria2Config.ACME.Dir
		for _, domain := range hysteria2Config.ACME.Domains {
			crtPath, err := util.FindFile(acmeDir, fmt.Sprintf("%s.crt", domain))
			if err != nil {
				continue
			}
			keyPath, err := util.FindFile(acmeDir, fmt.Sprintf("%s.key", domain))
			if err != nil {
				continue
			}
			hysteria2AcmePathVo.CrtPath = crtPath
			hysteria2AcmePathVo.KeyPath = keyPath
			return hysteria2AcmePathVo, nil
		}
	}
	return vo.Hysteria2AcmePathVo{}, errors.New("cert not found")
}

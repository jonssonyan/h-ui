package service

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/model/entity"
)

func UpdateConfig(key string, value string) error {
	if key == constant.Hysteria2Enable && value == "1" {
		config, err := dao.GetConfig("key = ?", constant.Hysteria2Config)
		if err != nil {
			return err
		}
		if config.Value == nil || *config.Value == "" {
			return errors.New("hysteria2 config is empty")
		}
	}
	return dao.UpdateConfig([]string{key}, map[string]interface{}{"value": value})
}

func GetConfig(key string) (entity.Config, error) {
	return dao.GetConfig("key = ?", key)
}

func ListConfig(keys []string) ([]entity.Config, error) {
	return dao.ListConfig("key in ?", keys)
}

func GetHysteria2Config() (bo.Hysteria2ServerConfig, error) {
	var serverConfig bo.Hysteria2ServerConfig
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Config)
	if err != nil {
		return serverConfig, err
	}
	if err = yaml.Unmarshal([]byte(*config.Value), &serverConfig); err != nil {
		return serverConfig, err
	}
	return serverConfig, nil
}

func UpdateHysteria2Config(hysteria2ServerConfig bo.Hysteria2ServerConfig) error {
	getConfig, err := dao.GetConfig("key = ?", constant.HUIWebPort)
	if err != nil {
		return err
	}
	hysteria2ServerConfig.Auth.Type = "http"
	hysteria2ServerConfig.Auth.HTTP.URL = fmt.Sprintf("http://127.0.0.1:%s/hui/hysteria2/auth", *getConfig.Value)
	hysteria2ServerConfig.Auth.HTTP.Insecure = true
	hysteria2ServerConfig.TrafficStats.Secret = ""
	config, err := yaml.Marshal(hysteria2ServerConfig)
	if err != nil {
		return err
	}
	return dao.UpdateConfig([]string{constant.Hysteria2Config}, map[string]interface{}{"value": string(config)})
}

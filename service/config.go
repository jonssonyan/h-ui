package service

import (
	"errors"
	"gopkg.in/yaml.v3"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/model/entity"
)

func UpdateConfig(key string, value string) error {
	if key == constant.Hysteria2Enable {
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

func ListConfig(keys []string) ([]entity.Config, error) {
	return dao.ListConfig("key in ?", keys)
}

func GetHysteria2Config() (bo.Hysteria2ServerConfig, error) {
	var serverConfig bo.Hysteria2ServerConfig
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Config)
	if err != nil {
		return serverConfig, err
	}
	if err := yaml.Unmarshal([]byte(*config.Value), &serverConfig); err != nil {
		return serverConfig, err
	}
	return serverConfig, nil
}

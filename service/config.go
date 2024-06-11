package service

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/model/entity"
	"strconv"
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
		// 启动Hysteria2
		if err = StartHysteria2(); err != nil {
			return err
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

func ListConfigNotIn(keys []string) ([]entity.Config, error) {
	return dao.ListConfig("key not in ?", keys)
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
	// 默认值
	getConfig, err := dao.GetConfig("key = ?", constant.HUIWebPort)
	if err != nil {
		return err
	}

	authType := "http"
	authHttpUrl := fmt.Sprintf("http://127.0.0.1:%s/hui/hysteria2/auth", *getConfig.Value)
	authHttpInsecure := true
	var auth bo.ServerConfigAuth
	auth.Type = &authType
	var http bo.ServerConfigAuthHTTP
	http.URL = &authHttpUrl
	http.Insecure = &authHttpInsecure
	auth.HTTP = &http
	hysteria2ServerConfig.Auth = &auth
	trafficStatsSecret := ""
	hysteria2ServerConfig.TrafficStats.Secret = &trafficStatsSecret

	config, err := yaml.Marshal(hysteria2ServerConfig)
	if err != nil {
		return err
	}
	return dao.UpdateConfig([]string{constant.Hysteria2Config}, map[string]interface{}{"value": string(config)})
}

func SetHysteria2Config(hysteria2ServerConfig bo.Hysteria2ServerConfig) error {
	config, err := yaml.Marshal(hysteria2ServerConfig)
	if err != nil {
		return err
	}
	return dao.UpdateConfig([]string{constant.Hysteria2Config}, map[string]interface{}{"value": string(config)})
}

func UpsertConfig(configs []entity.Config) error {
	return dao.UpsertConfig(configs)
}

func GetHysteria2ApiPort() (int64, error) {
	hysteria2Config, err := GetHysteria2Config()
	if err != nil {
		logrus.Errorf("get hysteria2 config err: %v", err)
		return 0, err
	}
	apiPort, err := strconv.ParseInt(*hysteria2Config.TrafficStats.Listen, 10, 64)
	if err != nil {
		logrus.Errorf("apiPort string conv int64 err: %v", err)
		return 0, err
	}
	return apiPort, nil
}

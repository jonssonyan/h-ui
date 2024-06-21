package service

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/proxy"
	"strconv"
	"strings"
	"time"
)

func Hysteria2Auth(conPass string) (string, error) {
	now := time.Now().UnixMilli()
	account, err := dao.GetAccount("deleted = 0 and con_pass = ? and ? < expire_time and ? > kick_util_time and (quota < 0 or quota > download + upload)", conPass, now, now)
	if err != nil {
		return "", err
	}

	// 限制设备数
	onlineUsers, err := Hysteria2Online()
	if err != nil {
		return "", err
	}
	device, exist := onlineUsers[*account.ConPass]
	if exist && *account.DeviceNo < device {
		return "", errors.New("device limited")
	}

	return *account.Username, nil
}

func Hysteria2Online() (map[string]int64, error) {
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Enable)
	if err != nil {
		return nil, err
	}
	if *config.Value == "1" {
		apiPort, err := GetHysteria2ApiPort()
		if err != nil {
			return nil, errors.New("get hysteria2 apiPort err")
		}
		hysteria2Api := proxy.NewHysteria2Api(apiPort)
		onlineUsers, err := hysteria2Api.OnlineUsers()
		if err != nil {
			return nil, err
		}
		return onlineUsers, nil
	}
	return map[string]int64{}, nil
}

func Hysteria2Kick(ids []int64, kickUtilTime int64) error {
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Enable)
	if err != nil {
		return err
	}
	if *config.Value != "1" {
		return errors.New("hysteria2 not enable")
	}
	if err := dao.UpdateAccount(ids, map[string]interface{}{"kick_util_time": kickUtilTime}); err != nil {
		return err
	}

	accounts, err := dao.ListAccount("id in ?", ids)
	if err != nil {
		return err
	}
	var keys []string
	for _, item := range accounts {
		keys = append(keys, *item.ConPass)
	}
	apiPort, err := GetHysteria2ApiPort()
	if err != nil {
		return errors.New("get hysteria2 apiPort err")
	}
	hysteria2Api := proxy.NewHysteria2Api(apiPort)
	if err = hysteria2Api.KickUsers(keys); err != nil {
		return err
	}
	return nil
}

func Hysteria2SubscribeUrl(accountId int64, protocol string, host string) (string, error) {
	account, err := dao.GetAccount("id = ?", accountId)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s//%s/hui/%s", protocol, host, *account.ConPass), nil
}

func Hysteria2Subscribe(conPass string, clientType string, host string) (string, string, error) {
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Config)
	if err != nil {
		return "", "", err
	}
	if config.Value == nil || *config.Value == "" {
		return "", "", errors.New("hysteria2 config is empty")
	}
	hysteria2Config, err := GetHysteria2Config()
	if err != nil {
		return "", "", err
	}

	account, err := dao.GetAccount("con_pass = ?", conPass)
	if err != nil {
		return "", "", err
	}

	userInfo := ""
	configYaml := ""
	if clientType == "Shadowrocket" || clientType == "ClashforWindows" {
		userInfo = fmt.Sprintf("upload=%d; download=%d; total=%d; expire=%d",
			*account.Upload,
			*account.Download,
			*account.Quota,
			*account.ExpireTime/1000)

		clashConfig := bo.ClashConfig{}
		var ClashConfigInterface []interface{}

		port, err := strconv.ParseUint(strings.TrimPrefix(*hysteria2Config.Listen, ":"), 10, 64)
		if err != nil {
			return "", "", err
		}

		hysteria2 := bo.Hysteria2{
			Name:     "hysteria2",
			Type:     "hysteria2",
			Server:   strings.Split(host, ":")[0],
			Port:     uint(port),
			Password: conPass,
		}
		if hysteria2Config.Obfs != nil &&
			hysteria2Config.Obfs.Type != nil &&
			*hysteria2Config.Obfs.Type == "salamander" &&
			hysteria2Config.Obfs.Salamander != nil &&
			hysteria2Config.Obfs.Salamander.Password != nil &&
			*hysteria2Config.Obfs.Salamander.Password != "" {
			hysteria2.Obfs = "salamander"
			hysteria2.ObfsPassword = *hysteria2Config.Obfs.Salamander.Password
		}

		if hysteria2Config.Bandwidth != nil {
			if hysteria2Config.Bandwidth.Up != nil &&
				*hysteria2Config.Bandwidth.Up != "" {
				up, err := strconv.ParseUint(strings.Split(*hysteria2Config.Bandwidth.Up, " ")[0], 10, 64)
				if err != nil {
					return "", "", err
				}

				hysteria2.Up = int(up)
			}
			if hysteria2Config.Bandwidth.Down != nil &&
				*hysteria2Config.Bandwidth.Down != "" {
				down, err := strconv.ParseUint(strings.Split(*hysteria2Config.Bandwidth.Down, " ")[0], 10, 64)
				if err != nil {
					return "", "", err
				}

				hysteria2.Down = int(down)
			}
		}

		ClashConfigInterface = append(ClashConfigInterface, hysteria2)

		proxyGroups := make([]bo.ProxyGroup, 0)
		proxyGroup := bo.ProxyGroup{
			Name:    "PROXY",
			Type:    "select",
			Proxies: []string{"hysteria2"},
		}

		proxyGroups = append(proxyGroups, proxyGroup)
		clashConfig.ProxyGroups = proxyGroups
		clashConfig.Proxies = ClashConfigInterface

		clashConfigYaml, err := yaml.Marshal(&clashConfig)
		if err != nil {
			return "", "", err
		}
		configYaml = string(clashConfigYaml)
	}

	return userInfo, configYaml, nil
}

func Hysteria2Url(accountId int64, hostname string) (string, error) {
	config, err := dao.GetConfig("key = ?", constant.Hysteria2Config)
	if err != nil {
		return "", err
	}
	if config.Value == nil || *config.Value == "" {
		return "", errors.New("hysteria2 config is empty")
	}
	hysteria2Config, err := GetHysteria2Config()
	if err != nil {
		return "", err
	}

	account, err := dao.GetAccount("id = ?", accountId)
	if err != nil {
		return "", err
	}

	urlConfig := ""
	if hysteria2Config.Obfs != nil &&
		hysteria2Config.Obfs.Type != nil &&
		*hysteria2Config.Obfs.Type == "salamander" &&
		hysteria2Config.Obfs.Salamander != nil &&
		hysteria2Config.Obfs.Salamander.Password != nil &&
		*hysteria2Config.Obfs.Salamander.Password != "" {
		urlConfig += fmt.Sprintf("&obfs=salamander&obfs-password=%s", *hysteria2Config.Obfs.Salamander.Password)
	}

	if urlConfig != "" {
		urlConfig = "?" + strings.TrimPrefix(urlConfig, "&")
	}

	return fmt.Sprintf("hysteria2://%s@%s%s", *account.ConPass, hostname, *hysteria2Config.Listen) + urlConfig, nil
}

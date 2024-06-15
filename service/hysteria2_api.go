package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/model/vo"
	"h-ui/proxy"
	"strings"
)

func Hysteria2Auth(conPass string) (string, error) {
	account, err := dao.GetAccount("deleted = 0 and pass = ? and CURRENT_TIMESTAMP < expire_time and CURRENT_TIMESTAMP > kick_util_time and (quota < 0 or quota > download + upload) ", conPass)
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

func Hysteria2Url(c *gin.Context) (vo.Hysteria2UrlVo, error) {
	hysteria2UrlVo := vo.Hysteria2UrlVo{}
	accountInfo, err := GetAccountInfo(c)
	if err != nil {
		return hysteria2UrlVo, err
	}
	hysteria2Config, err := GetHysteria2Config()
	if err != nil {
		return hysteria2UrlVo, err
	}

	account, err := dao.GetAccount("id = ?", accountInfo.Id)
	if err != nil {
		return hysteria2UrlVo, err
	}

	hysteria2UrlVo.Auth = *account.ConPass

	config := ""

	if hysteria2Config.Obfs != nil &&
		hysteria2Config.Obfs.Type != nil &&
		*hysteria2Config.Obfs.Type == "salamander" &&
		hysteria2Config.Obfs.Salamander != nil &&
		hysteria2Config.Obfs.Salamander.Password != nil &&
		*hysteria2Config.Obfs.Salamander.Password != "" {
		config += fmt.Sprintf("&obfs=salamander&obfs-password=%s", *hysteria2Config.Obfs.Salamander.Password)
	}

	hysteria2UrlVo.Config = strings.TrimPrefix(config, "&")

	return hysteria2UrlVo, nil
}

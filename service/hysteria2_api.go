package service

import (
	"errors"
	"h-ui/dao"
	"h-ui/proxy"
)

func Hysteria2Auth(conPass string) (string, error) {
	account, err := dao.GetAccount("deleted = 0 and pass = ? and CURRENT_TIMESTAMP < expire_time and (quota < 0 or quota > download + upload) ", conPass)
	if err != nil {
		return "", err
	}
	return *account.Username, nil
}

func Hysteria2Online() (map[string]int64, error) {
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

func Hysteria2Kick(usernames []string) error {
	accounts, err := dao.ListAccount("username in ?", usernames)
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

package service

import (
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/proxy"
	"h-ui/util"
	"strconv"
)

func CronHandleAccount() {
	go func() {
		hysteriaEnable, err := dao.GetConfig("key = ?", constant.Hysteria2Enable)
		if err != nil {
			logrus.Errorf("get hysteria2 enable config err: %v", err)
			return
		}
		if hysteriaEnable.Value != nil && *hysteriaEnable.Value == "1" {
			hysteria2Config, err := GetHysteria2Config()
			if err != nil {
				logrus.Errorf("get hysteria2 config err: %v", err)
				return
			}
			apiPort, err := strconv.ParseInt(hysteria2Config.TrafficStats.Listen, 10, 64)
			if err != nil {
				logrus.Errorf("apiPort string conv int64 err: %v", err)
				return
			}

			hysteria2TrafficTime, err := dao.GetConfig("key = ?", constant.Hysteria2TrafficTime)
			if err != nil {
				logrus.Errorf("get hysteria2 traffic time err: %v", err)
				return
			}
			hysteria2TrafficTimeInt, err := strconv.ParseInt(*hysteria2TrafficTime.Value, 10, 64)
			if err != nil {
				logrus.Errorf("hysteria2TrafficTime string conv int64 err: %v", err)
				return
			}

			hysteria2Api := proxy.NewHysteria2Api(apiPort)
			users, err := hysteria2Api.ListUsers(true)
			if err == nil {
				userLists := util.SplitMap(users, 10)
				for _, userList := range userLists {
					go func(users map[string]bo.Hysteria2UserTraffic) {
						for conPass, traffic := range users {
							conPassEncrypt := util.SHA224String(conPass)
							if err = dao.UpdateAccountTraffic(conPassEncrypt, traffic.Rx*hysteria2TrafficTimeInt, traffic.Tx*hysteria2TrafficTimeInt); err != nil {
								logrus.Errorf("hysteria2 hanle account err apiPort: %d err: %v", apiPort, err)
								continue
							}
						}
					}(userList)
				}
			}
		}
	}()
}

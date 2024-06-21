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
			apiPort, err := GetHysteria2ApiPort()
			if err != nil {
				logrus.Errorf("get hysteria2 apiPort err: %v", err)
				return
			}

			hysteria2TrafficTime, err := dao.GetConfig("key = ?", constant.Hysteria2TrafficTime)
			if err != nil {
				logrus.Errorf("get hysteria2 traffic time err: %v", err)
				return
			}
			hysteria2TrafficTimeFloat, err := strconv.ParseFloat(*hysteria2TrafficTime.Value, 64)
			if err != nil {
				logrus.Errorf("hysteria2TrafficTime string conv int64 err: %v", err)
				return
			}

			users, err := proxy.NewHysteria2Api(apiPort).ListUsers(true)
			if err == nil {
				userLists := util.SplitMap(users, 10)
				for _, userList := range userLists {
					go func(users map[string]bo.Hysteria2UserTraffic) {
						for conPass, traffic := range users {
							conPassEncrypt := util.SHA224String(conPass)
							if err = dao.UpdateAccountTraffic(conPassEncrypt, int64(float64(traffic.Rx)*hysteria2TrafficTimeFloat), int64(float64(traffic.Tx)*hysteria2TrafficTimeFloat)); err != nil {
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

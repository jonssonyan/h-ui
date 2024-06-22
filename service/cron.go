package service

import (
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/proxy"
	"h-ui/util"
	"strconv"
	"time"
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

			jwtSecretConfig, err := dao.GetConfig("key = ?", constant.JwtSecret)
			if err != nil {
				logrus.Errorf("jwtSecretConfig is nill")
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

			// 保存流量数据
			go saveAccountTraffic(apiPort, *jwtSecretConfig.Value, hysteria2TrafficTimeFloat)
			// 踢下线
			go kickAccount(apiPort, *jwtSecretConfig.Value)
		}
	}()
}

func saveAccountTraffic(apiPort int64, jwtSecret string, hysteria2TrafficTimeFloat float64) {
	users, err := proxy.NewHysteria2Api(apiPort).ListUsers(true, jwtSecret)
	if err != nil {
		logrus.Errorf("saveAccountTraffic ListUsers err: %v", err)
		return
	}
	if len(users) > 0 {
		userLists := util.SplitMap(users, 10)
		for _, userList := range userLists {
			go func(userList map[string]bo.Hysteria2UserTraffic) {
				for username, traffic := range userList {
					if err = dao.UpdateAccountTraffic(username, int64(float64(traffic.Rx)*hysteria2TrafficTimeFloat), int64(float64(traffic.Tx)*hysteria2TrafficTimeFloat)); err != nil {
						logrus.Errorf("hysteria2 save account traffic err apiPort: %d err: %v", apiPort, err)
						continue
					}
				}
			}(userList)
		}
	}
}

func kickAccount(apiPort int64, jwtSecret string) {
	users, err := proxy.NewHysteria2Api(apiPort).OnlineUsers(jwtSecret)
	if err != nil {
		logrus.Errorf("kickAccount OnlineUsers err: %v", err)
		return
	}
	if len(users) > 0 {
		i := 0
		usernames := make([]string, len(users))
		for k := range users {
			usernames[i] = k
			i++
		}
		usernameLists := util.SplitArr(usernames, 10)
		for _, usernameList := range usernameLists {
			go func(usernameList []string) {
				now := time.Now().UnixMilli()
				accounts, err := dao.ListAccount("username in ? and (deleted = 1 or (quota > 0 and quota < download + upload)) or ? > expire_time or ? < kick_util_time", usernameList, now, now)
				if err != nil {
					logrus.Errorf(err.Error())
					return
				}
				kickUsernames := make([]string, len(accounts))
				j := 0
				for _, item := range accounts {
					kickUsernames[j] = *item.Username
					j++
				}
				if err = proxy.NewHysteria2Api(apiPort).KickUsers(kickUsernames, jwtSecret); err != nil {
					logrus.Errorf("hysteria2 kick account err apiPort: %d err: %v", apiPort, err)
					return
				}
			}(usernameList)
		}
	}
}

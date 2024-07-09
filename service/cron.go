package service

import (
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/bo"
	"h-ui/model/constant"
	"h-ui/proxy"
	"h-ui/util"
	"strconv"
	"sync"
	"time"
)

var trafficMutex sync.Mutex
var kickMutex sync.Mutex

func CronHandleAccount() {
	go func() {
		hysteriaEnable, err := dao.GetConfig("key = ?", constant.Hysteria2Enable)
		if err != nil {
			return
		}
		if hysteriaEnable.Value != nil && *hysteriaEnable.Value == "1" {
			apiPort, err := GetHysteria2ApiPort()
			if err != nil {
				return
			}

			jwtSecretConfig, err := dao.GetConfig("key = ?", constant.JwtSecret)
			if err != nil {
				return
			}

			// 保存流量数据
			go saveAccountTraffic(apiPort, *jwtSecretConfig.Value)

			// 踢下线
			go kickAccount(apiPort, *jwtSecretConfig.Value)
		}
	}()
}

func saveAccountTraffic(apiPort int64, jwtSecret string) {
	if !trafficMutex.TryLock() {
		return
	}
	defer trafficMutex.Unlock()

	hysteria2TrafficTime, err := dao.GetConfig("key = ?", constant.Hysteria2TrafficTime)
	if err != nil {
		return
	}
	hysteria2TrafficTimeFloat, err := strconv.ParseFloat(*hysteria2TrafficTime.Value, 64)
	if err != nil {
		logrus.Errorf("hysteria2TrafficTime string conv int64 err: %v", err)
		return
	}

	users, err := proxy.NewHysteria2Api(apiPort).ListUsers(true, jwtSecret)
	if err != nil {
		return
	}
	if len(users) > 0 {
		userLists := util.SplitMap(users, 10)
		var wg sync.WaitGroup
		for _, userList := range userLists {
			wg.Add(1)
			go func(userList map[string]bo.Hysteria2UserTraffic) {
				defer wg.Done()
				for username, traffic := range userList {
					if err = dao.UpdateAccountTraffic(username, int64(float64(traffic.Rx)*hysteria2TrafficTimeFloat), int64(float64(traffic.Tx)*hysteria2TrafficTimeFloat)); err != nil {
						continue
					}
				}
			}(userList)
		}
		wg.Wait()
	}
}

func kickAccount(apiPort int64, jwtSecret string) {
	if !kickMutex.TryLock() {
		return
	}
	defer kickMutex.Unlock()

	users, err := proxy.NewHysteria2Api(apiPort).OnlineUsers(jwtSecret)
	if err != nil {
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
		var wg sync.WaitGroup
		for _, usernameList := range usernameLists {
			wg.Add(1)
			go func(usernameList []string) {
				defer wg.Done()
				now := time.Now().UnixMilli()
				accounts, err := dao.ListAccount("username in ? and (deleted = 1 or (quota > 0 and quota < download + upload)) or ? > expire_time or ? < kick_util_time", usernameList, now, now)
				if err != nil {
					return
				}
				kickUsernames := make([]string, len(accounts))
				j := 0
				for _, item := range accounts {
					kickUsernames[j] = *item.Username
					j++
				}
				if err = proxy.NewHysteria2Api(apiPort).KickUsers(kickUsernames, jwtSecret); err != nil {
					return
				}
			}(usernameList)
		}
		wg.Wait()
	}
}

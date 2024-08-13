package middleware

import (
	"errors"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/model/constant"
	"h-ui/service"
	"time"
)

func InitCron() error {
	loc := time.Now().Location()
	c := cron.New(cron.WithLocation(loc))
	_, err := c.AddFunc("@every 30s", service.CronHandleAccount)
	if err != nil {
		logrus.Errorf("cron add func CronHandleAccount err: %v", err)
		return errors.New("cron add func CronHandleAccount err")
	}
	resetTrafficCron, err := dao.GetConfig("key = ?", constant.ResetTrafficCron)
	if err != nil {
		return err
	}
	if *resetTrafficCron.Value != "" {
		_, err := c.AddFunc(*resetTrafficCron.Value, service.CronResetTraffic)
		if err != nil {
			logrus.Errorf("cron add func CronResetTraffic err: %v", err)
		}
	}
	c.Start()
	return nil
}

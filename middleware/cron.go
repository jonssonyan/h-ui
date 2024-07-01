package middleware

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"h-ui/service"
	"time"
)

func InitCron() error {
	loc := time.Now().Location()
	c := cron.New(cron.WithLocation(loc))
	_, err := c.AddFunc("@every 30s", service.CronHandleAccount)
	if err != nil {
		logrus.Errorf("cron add func err: %v", err)
		return err
	}
	c.Start()
	return nil
}

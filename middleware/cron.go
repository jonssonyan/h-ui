package middleware

import (
	"github.com/robfig/cron/v3"
	"h-ui/service"
	"time"
)

func InitCron() {
	loc := time.Now().Location()
	c := cron.New(cron.WithLocation(loc))
	_, _ = c.AddFunc("@every 30s", service.CronHandleAccount)
	c.Start()
}

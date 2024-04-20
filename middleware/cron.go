package middleware

import (
	"github.com/robfig/cron/v3"
	"h-ui/service"
	"time"
)

func InitCron() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(location))
	_, _ = c.AddFunc("@every 30s", service.CronHandlerAccount)
	c.Start()
}

package middleware

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"h-ui/service"
	"time"
)

func InitCron() error {
	loc := time.Now().Location()
	c := cron.New(cron.WithLocation(loc))
	_, err := c.AddFunc("@every 30s", service.CronHandleAccount)
	if err != nil {
		return errors.New(fmt.Sprintf("cron add func err: %v", err))
	}
	c.Start()
	return nil
}

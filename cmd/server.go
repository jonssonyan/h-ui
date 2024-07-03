package cmd

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/middleware"
	"h-ui/model/constant"
	"h-ui/router"
	"h-ui/service"
	"h-ui/util"
	"net/http"
	"os"
)

func runServer() error {
	defer releaseResource()

	middleware.InitLog()
	if err := initFile(); err != nil {
		logrus.Errorf("init file err: %v", err)
		return err
	}
	if err := dao.InitSqliteDB(port); err != nil {
		logrus.Errorf("init sqlite db err: %v", err)
		return err
	}
	if err := middleware.InitCron(); err != nil {
		logrus.Errorf("init cron err: %v", err)
		return err
	}
	if err := service.InitHysteria2(); err != nil {
		logrus.Errorf("init Hysteria2 err: %v", err)
		return err
	}

	r := gin.Default()
	router.Router(r)

	port, crtPath, keyPath, err := service.GetServerPortAndCert()
	if err != nil {
		logrus.Errorf("get server port and cert err: %v", err)
		return err
	}

	service.InitServer(fmt.Sprintf(":%d", port), r)
	if err := service.StartServer(crtPath, keyPath); err != nil && err != http.ErrServerClosed {
		logrus.Errorf("start server err: %v", err)
		return err
	}
	return nil
}

func releaseResource() {
	if err := dao.CloseSqliteDB(); err != nil {
		logrus.Errorf(err.Error())
	}
	if err := service.ReleaseHysteria2(); err != nil {
		logrus.Errorf(err.Error())
	}
}

func initFile() error {
	var dirs = []string{constant.LogDir, constant.SqliteDBDir, constant.BinDir, constant.ExportPathDir}
	for _, item := range dirs {
		if !util.Exists(item) {
			if err := os.Mkdir(item, os.ModePerm); err != nil {
				return errors.New(fmt.Sprintf("%s create err: %v", item, err))
			}
		}
	}
	return nil
}

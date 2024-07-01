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

	if err := initFile(); err != nil {
		return err
	}
	middleware.InitLog()
	if err := dao.InitSqliteDB(port); err != nil {
		return err
	}
	if err := middleware.InitCron(); err != nil {
		return err
	}
	if err := service.InitHysteria2(); err != nil {
		return err
	}

	r := gin.Default()
	router.Router(r)

	webServer, err := service.NewServer()
	if err != nil {
		return err
	}
	if err = webServer.StartServer(r); err != nil && err != http.ErrServerClosed {
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
				logrus.Errorf("%s create err: %v", item, err)
				return errors.New(fmt.Sprintf("%s create err: %v", item, err))
			}
		}
	}
	return nil
}

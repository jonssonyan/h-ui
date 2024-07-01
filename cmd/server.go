package cmd

import (
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

	initFile()
	middleware.InitLog()
	dao.InitSqliteDB(port)
	middleware.InitCron()
	service.InitHysteria2()

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

func initFile() {
	var dirs = []string{constant.LogDir, constant.SqliteDBDir, constant.BinDir, constant.ExportPathDir}
	for _, item := range dirs {
		if !util.Exists(item) {
			if err := os.Mkdir(item, os.ModePerm); err != nil {
				panic(fmt.Sprintf("%s create err: %v", item, err))
			}
		}
	}
}

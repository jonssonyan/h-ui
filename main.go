package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"h-ui/controller"
	"h-ui/dao"
	"h-ui/middleware"
	"h-ui/model/constant"
	"h-ui/router"
	"h-ui/util"
	"os"
)

func main() {
	config, err := dao.GetConfig("key = ?", constant.WebServerPort)
	if err != nil {
		logrus.Errorf("webServerPort get err: %v", err)
		panic(err)
	}
	r := gin.Default()
	router.Router(r)
	_ = r.Run(fmt.Sprintf(":%s", *config.Value))
	defer releaseResource()
}

func init() {
	initFile()
	middleware.InitLog()
	dao.InitSqliteDB()
	middleware.InitCron()
	controller.InitValidator()
}

func releaseResource() {
	dao.CloseSqliteDB()
}

func initFile() {

	var dirs = []string{constant.LogDir, constant.ProxyDir, constant.SqliteDBDir}
	for _, item := range dirs {
		if !util.Exists(item) {
			if err := os.Mkdir(item, os.ModePerm); err != nil {
				logrus.Errorf("%s create err: %v", item, err)
				panic(err)
			}
		}
	}

	var files = []string{constant.SqliteDBPath}
	for _, item := range files {
		if !util.Exists(item) {
			func() {
				file, err := os.Create(item)
				if err != nil {
					logrus.Errorf("%s file create err: %v", item, err)
					panic(err)
				}
				defer file.Close()
			}()
		}
	}
}

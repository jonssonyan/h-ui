package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"h-ui/dao"
	"h-ui/middleware"
	"h-ui/model/constant"
	"h-ui/router"
	"h-ui/service"
	"h-ui/util"
	"os"
)

func main() {
	defer releaseResource()
	r := gin.Default()
	router.Router(r)
	if err := service.StartServer(r); err != nil {
		panic(err)
	}
}

func init() {
	initFile()
	middleware.InitLog()
	dao.InitSqliteDB()
	middleware.InitCron()
	service.InitHysteria2()
}

func releaseResource() {
	dao.CloseSqliteDB()
}

func initFile() {

	var dirs = []string{constant.LogDir, constant.SqliteDBDir, constant.BinDir, constant.ExportPathDir}
	for _, item := range dirs {
		if !util.Exists(item) {
			if err := os.Mkdir(item, os.ModePerm); err != nil {
				logrus.Errorf("%s create err: %v", item, err)
				panic(err)
			}
		}
	}
}

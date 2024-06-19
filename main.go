package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"h-ui/cmd"
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
	for {
		if err := service.StartServer(r); err != nil {
			panic(err)
		}
	}
}

func init() {
	cmd.InitCmd()
	initFile()
	middleware.InitLog()
	dao.InitSqliteDB()
	middleware.InitCron()
	service.InitHysteria2()
}

func releaseResource() {
	dao.CloseSqliteDB()
	service.ReleaseHysteria2()
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

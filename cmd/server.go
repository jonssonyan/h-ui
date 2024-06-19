package cmd

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"h-ui/dao"
	"h-ui/middleware"
	"h-ui/model/constant"
	"h-ui/router"
	"h-ui/service"
	"h-ui/util"
	"os"
	"strconv"
)

var port string

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

func runServer() {
	initFile()
	middleware.InitLog()
	dao.InitSqliteDB(port)
	middleware.InitCron()
	service.InitHysteria2()
	defer releaseResource()
	r := gin.Default()
	router.Router(r)
	for {
		if err := service.StartServer(r); err != nil {
			panic(err)
		}
	}
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start web server",
	Long:  "Start web server",
	Run:   runServerCmd,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "Set the port for the web server")
	rootCmd.AddCommand(serverCmd)
}

func runServerCmd(cmd *cobra.Command, args []string) {
	if err := validateAndSetPort(port); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	runServer()
}

func validateAndSetPort(p string) error {
	if p != "" {
		value, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			return errors.New("invalid port value")
		}
		if value <= 0 || value > 65535 {
			return errors.New("the port range is between 0-65535")
		}
	}
	return nil
}

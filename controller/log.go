package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/vo"
	"os"
	"strings"
)

func SystemLog(c *gin.Context) {
	logData, err := os.ReadFile("logs/h-ui.log")
	if err != nil {
		vo.Fail("Unable to read log file", c)
		return
	}
	logLines := strings.Split(string(logData), "\n")
	vo.Success(logLines, c)
}

func HysteriaLog2(c *gin.Context) {
	logData, err := os.ReadFile("logs/hysteria2.log")
	if err != nil {
		vo.Fail("Unable to read log file", c)
		return
	}
	logLines := strings.Split(string(logData), "\n")
	vo.Success(logLines, c)
}

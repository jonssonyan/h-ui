package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/vo"
	"h-ui/service"
)

func MonitorSystem(c *gin.Context) {
	systemMonitorVo, err := service.MonitorSystem()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(systemMonitorVo, c)
}

func MonitorHysteria2(c *gin.Context) {
	hysteria2MonitorVo, err := service.MonitorHysteria2()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(hysteria2MonitorVo, c)
}

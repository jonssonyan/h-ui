package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/vo"
	"h-ui/util"
)

func MonitorSystem(c *gin.Context) {
	cpuPercent, err := util.GetCpuPercent()
	if err != nil {
		vo.Fail("cpu query failed", c)
		return
	}
	memPercent, err := util.GetMemPercent()
	if err != nil {
		vo.Fail("mem query failed", c)
		return
	}
	diskPercent, err := util.GetDiskPercent()
	if err != nil {
		vo.Fail("disk query failed", c)
		return
	}
	vo.Success(vo.SystemMonitorVo{
		CpuPercent:  cpuPercent,
		MemPercent:  memPercent,
		DiskPercent: diskPercent,
	}, c)
}

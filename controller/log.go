package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"h-ui/model/vo"
	"h-ui/util"
)

func LogSystem(c *gin.Context) {
	exists := util.Exists("logs/h-ui.log")
	logSystemVos := make([]vo.LogSystemVo, 0)
	if !exists {
		vo.Success(logSystemVos, c)
		return
	}

	logLines, err := util.ReadLinesFromBottom("logs/h-ui.log")
	if err != nil {
		vo.Fail("Unable to read log file", c)
		return
	}

	for _, line := range logLines {
		if line == "" {
			continue
		}
		logSystemVo := vo.LogSystemVo{}
		err := json.Unmarshal([]byte(line), &logSystemVo)
		if err != nil {
			vo.Fail("Unable to unmarshal log data", c)
			continue
		}
		logSystemVos = append(logSystemVos, logSystemVo)
	}
	vo.Success(logSystemVos, c)
}

func LogHysteria2(c *gin.Context) {
	logLines, err := util.ReadLinesFromBottom("logs/h-ui.log")
	if err != nil {
		vo.Fail("Unable to read log file", c)
		return
	}
	vo.Success(logLines, c)
}

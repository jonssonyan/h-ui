package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"h-ui/model/dto"
	"h-ui/model/vo"
	"h-ui/util"
)

func LogSystem(c *gin.Context) {
	logSystemDto, err := validateField(c, dto.LogSystemDto{})
	if err != nil {
		return
	}
	exists := util.Exists("logs/h-ui.log")
	logSystemVos := make([]vo.LogSystemVo, 0)
	if !exists {
		vo.Success(logSystemVos, c)
		return
	}
	numLine := 0
	if logSystemDto.NumLine != nil || *logSystemDto.NumLine > 0 {
		numLine = *logSystemDto.NumLine
	}
	logLines, total, err := util.ReadLinesFromBottom("logs/h-ui.log", numLine)
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
	vo.Success(vo.LogSystemPage{
		LogSystemVos: logSystemVos,
		Total:        int64(total),
	}, c)
}

func LogHysteria2(c *gin.Context) {
	vo.Success(nil, c)
}

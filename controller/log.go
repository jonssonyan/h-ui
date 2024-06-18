package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"h-ui/model/constant"
	"h-ui/model/dto"
	"h-ui/model/vo"
	"h-ui/util"
	"time"
)

func LogSystem(c *gin.Context) {
	logSystemDto, err := validateField(c, dto.LogDto{})
	if err != nil {
		return
	}
	exists := util.Exists(constant.SystemLogPath)
	logSystemVos := make([]vo.LogSystemVo, 0)
	if !exists {
		vo.Success(logSystemVos, c)
		return
	}
	numLine := 0
	if logSystemDto.NumLine != nil || *logSystemDto.NumLine > 0 {
		numLine = *logSystemDto.NumLine
	}
	logLines, total, err := util.ReadLinesFromBottom(constant.SystemLogPath, numLine)
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
	vo.Success(vo.LogSystemPage[vo.LogSystemVo]{
		LogSystemVos: logSystemVos,
		Total:        int64(total),
	}, c)
}

func LogHysteria2(c *gin.Context) {
	logSystemDto, err := validateField(c, dto.LogDto{})
	if err != nil {
		return
	}
	exists := util.Exists(constant.Hysteria2LogPath)
	logHysteria2Vos := make([]vo.LogHysteria2Vo, 0)
	if !exists {
		vo.Success(logHysteria2Vos, c)
		return
	}
	numLine := 0
	if logSystemDto.NumLine != nil || *logSystemDto.NumLine > 0 {
		numLine = *logSystemDto.NumLine
	}
	logLines, total, err := util.ReadLinesFromBottom(constant.Hysteria2LogPath, numLine)
	if err != nil {
		vo.Fail("Unable to read log file", c)
		return
	}

	for _, line := range logLines {
		if line == "" {
			continue
		}
		logHysteria2Vo := vo.LogHysteria2Vo{}
		err := json.Unmarshal([]byte(line), &logHysteria2Vo)
		if err != nil {
			vo.Fail("Unable to unmarshal log data", c)
			continue
		}
		logHysteria2Vos = append(logHysteria2Vos, logHysteria2Vo)
	}
	vo.Success(vo.LogSystemPage[vo.LogHysteria2Vo]{
		LogSystemVos: logHysteria2Vos,
		Total:        int64(total),
	}, c)
}

func ExportLog(c *gin.Context) {
	logExportDto, err := validateField(c, dto.LogExportDto{})
	if err != nil {
		return
	}

	var fileName string
	var filePath string
	if *logExportDto.Option == 0 {
		fileName = fmt.Sprintf("h-ui-%s.log", time.Now().Format("20060102150405"))
		filePath = constant.SystemLogPath
	} else if *logExportDto.Option == 1 {
		fileName = fmt.Sprintf("hysteria2-%s.log", time.Now().Format("20060102150405"))
		filePath = constant.Hysteria2LogPath
	}

	if !util.Exists(filePath) {
		vo.Fail("log file not exist", c)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.File(filePath)
}

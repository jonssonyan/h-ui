package service

import (
	"errors"
	"fmt"
	"h-ui/model/constant"
	"h-ui/model/vo"
	"h-ui/util"
	"regexp"
	"strings"
)

func MonitorSystem() (vo.SystemMonitorVo, error) {
	cpuPercent, err := util.GetCpuPercent()
	if err != nil {
		return vo.SystemMonitorVo{}, errors.New("cpu query failed")
	}
	memPercent, err := util.GetMemPercent()
	if err != nil {
		return vo.SystemMonitorVo{}, errors.New("mem query failed")
	}
	diskPercent, err := util.GetDiskPercent()
	if err != nil {
		return vo.SystemMonitorVo{}, errors.New("disk query failed")
	}
	return vo.SystemMonitorVo{
		HUIVersion:  constant.Version,
		CpuPercent:  cpuPercent,
		MemPercent:  memPercent,
		DiskPercent: diskPercent,
	}, nil
}

func MonitorHysteria2() (vo.Hysteria2MonitorVo, error) {
	var hysteria2MonitorVo vo.Hysteria2MonitorVo
	onlineUsers, err := Hysteria2Online()
	if err != nil {
		return hysteria2MonitorVo, err
	}

	if len(onlineUsers) > 0 {
		hysteria2MonitorVo.UserTotal = int64(len(onlineUsers))
		var deviceTotal int64 = 0
		for _, value := range onlineUsers {
			deviceTotal += value
		}
		hysteria2MonitorVo.DeviceTotal = deviceTotal
	}

	hysteria2MonitorVo.Version = "-"
	content, err := util.Exec(fmt.Sprintf("%s version", util.GetHysteria2BinPath()))
	if err == nil {
		pattern := `v\d+\.\d+\.\d+`
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(strings.TrimSpace(content), -1)
		if len(matches) > 0 {
			hysteria2MonitorVo.Version = matches[0]
		}
	}

	running := Hysteria2IsRunning()
	hysteria2MonitorVo.Running = running
	return hysteria2MonitorVo, nil
}

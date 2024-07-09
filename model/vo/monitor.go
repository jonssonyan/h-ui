package vo

type SystemMonitorVo struct {
	HUIVersion  string  `json:"huiVersion"`
	CpuPercent  float64 `json:"cpuPercent"`
	MemPercent  float64 `json:"memPercent"`
	DiskPercent float64 `json:"diskPercent"`
}

type Hysteria2MonitorVo struct {
	UserTotal   int64  `json:"userTotal"`   // 在线用户数
	DeviceTotal int64  `json:"deviceTotal"` // 在线设备数
	Version     string `json:"version"`     // 版本
	Running     bool   `json:"running"`     // 运行状态
}

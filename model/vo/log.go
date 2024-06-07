package vo

type LogSystemPage[T LogSystemVo | LogHysteria2Vo] struct {
	LogSystemVos []T   `json:"records"`
	Total        int64 `json:"total"`
}

type LogSystemVo struct {
	ClientIP    string `json:"clientIp"`
	LatencyTime int64  `json:"latencyTime"`
	Level       string `json:"level"`
	Msg         string `json:"msg"`
	ReqMethod   string `json:"reqMethod"`
	ReqUri      string `json:"reqUri"`
	StatusCode  int64  `json:"statusCode"`
	Time        string `json:"time"`
}

type LogHysteria2Vo struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}

package vo

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
}

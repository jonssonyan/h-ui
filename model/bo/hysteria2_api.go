package bo

type Hysteria2UserTraffic struct {
	Tx int64 `json:"tx"` // upload
	Rx int64 `json:"rx"` // download
}

type Hysteria2User struct {
	Pass string `json:"Pass"`
	Tx   int64  `json:"tx"` // upload
	Rx   int64  `json:"rx"` // download
}

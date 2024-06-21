package bo

type Hysteria2UserTraffic struct {
	Tx int64 `json:"tx"` // upload
	Rx int64 `json:"rx"` // download
}

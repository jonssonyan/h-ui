package vo

type AccountVo struct {
	BaseVo
	Username     string `json:"username"`
	Quota        int64  `json:"quota"`
	Download     int64  `json:"download"`
	Upload       int64  `json:"upload"`
	ExpireTime   int64  `json:"expireTime"`
	KickUtilTime int64  `json:"kickUtilTime"`
	DeviceNo     int64  `json:"deviceNo"` // 限制设备数
	Role         string `json:"role"`
	Deleted      int64  `json:"deleted"`

	Online bool  `json:"online"` // 是否在线
	Device int64 `json:"device"` // 在线设备数
}
type AccountPageVo struct {
	AccountVos []AccountVo `json:"records"`
	Total      int64       `json:"total"`
}

type AccountInfoVo struct {
	Id       int64    `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

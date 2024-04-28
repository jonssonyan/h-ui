package vo

type AccountVo struct {
	BaseVo
	Username   string `json:"username"`
	Quota      int64  `json:"quota"`
	Download   int64  `json:"download"`
	Upload     int64  `json:"upload"`
	ExpireTime int64  `json:"expireTime"`
	Deleted    int64  `json:"deleted"`
}
type AccountPageVo struct {
	AccountVos []AccountVo `json:"records"`
	Total      int64       `json:"total"`
}

type AccountInfoVo struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	IsAdmin  int64  `json:"isAdmin"`
}

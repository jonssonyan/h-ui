package bo

import "time"

type AccountBo struct {
	Id       int64    `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Deleted  int64    `json:"deleted"`
}

type AccountExport struct {
	Id           int64     `json:"id"`
	Username     string    `json:"username"`
	Pass         string    `json:"pass"`
	ConPass      string    `json:"conPass"`
	Quota        int64     `json:"quota"`
	Download     int64     `json:"download"`
	Upload       int64     `json:"upload"`
	ExpireTime   int64     `json:"expireTime"`
	DeviceNo     int64     `json:"deviceNo"`
	KickUtilTime int64     `json:"kickUtilTime"`
	Role         string    `json:"role"`
	Deleted      int64     `json:"deleted"`
	CreateTime   time.Time `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
}

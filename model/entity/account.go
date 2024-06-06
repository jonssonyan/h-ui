package entity

type Account struct {
	Username     *string `gorm:"column:username;default:''" json:"username"`
	Pass         *string `gorm:"column:pass;default:''" json:"pass"`
	ConPass      *string `gorm:"column:con_pass;default:''" json:"conPass"`
	Quota        *int64  `gorm:"column:quota;default:0" json:"quota"`
	Download     *int64  `gorm:"column:download;default:0" json:"download"`
	Upload       *int64  `gorm:"column:upload;default:0" json:"upload"`
	ExpireTime   *int64  `gorm:"column:expire_time;default:0" json:"expireTime"`
	KickUtilTime *int64  `gorm:"column:kick_util_time;default:0" json:"kickUtilTime"`
	DeviceNo     *int64  `gorm:"column:device_no;default:3" json:"deviceNo"`
	Role         *string `gorm:"column:role;default:'user'" json:"role"`
	Deleted      *int64  `gorm:"column:deleted;default:0" json:"deleted"`
	BaseEntity   `gorm:"embedded"`
}

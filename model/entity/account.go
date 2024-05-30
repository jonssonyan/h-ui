package entity

type Account struct {
	Username     *string `gorm:"column:username;default:''"`
	Pass         *string `gorm:"column:pass;default:''"`
	ConPass      *string `gorm:"column:con_pass;default:''"`
	Quota        *int64  `gorm:"column:quota;default:0"`
	Download     *int64  `gorm:"column:download;default:0"`
	Upload       *int64  `gorm:"column:upload;default:0"`
	ExpireTime   *int64  `gorm:"column:expire_time;default:0"`
	KickUtilTime *int64  `gorm:"column:kick_util_time;default:0"`
	DeviceNo     *int64  `gorm:"column:device_no;default:3"`
	Role         *string `gorm:"column:role;default:'user'"`
	Deleted      *int64  `gorm:"column:deleted;default:0"`
	BaseEntity   `gorm:"embedded"`
}

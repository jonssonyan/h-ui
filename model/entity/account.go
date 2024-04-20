package entity

type Account struct {
	Username   *string `gorm:"column:username;default:''"`
	Pass       *string `gorm:"column:pass;default:''"`
	Quota      *int64  `gorm:"column:quota;default:0"`
	Download   *int64  `gorm:"column:download;default:0"`
	Upload     *int64  `gorm:"column:upload;default:0"`
	ExpireTime *int64  `gorm:"column:expire_time;default:0"`
	Deleted    *int64  `gorm:"column:deleted;default:0"`
	BaseEntity `gorm:"embedded"`
}

package entity

type Config struct {
	Key        *string `gorm:"column:key;default:''"`
	Value      *string `gorm:"column:value;default:''"`
	Remark     *string `gorm:"column:remark;default:''"`
	BaseEntity `gorm:"embedded"`
}

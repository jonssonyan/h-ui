package entity

type Config struct {
	Key        *string `gorm:"column:key;default:''" json:"key"`
	Value      *string `gorm:"column:value;default:''" json:"value"`
	Remark     *string `gorm:"column:remark;default:''" json:"remark"`
	BaseEntity `gorm:"embedded"`
}

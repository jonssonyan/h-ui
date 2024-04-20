package entity

import "time"

type BaseEntity struct {
	Id         *int64     `gorm:"column:id;primaryKey"`
	CreateTime *time.Time `gorm:"column:create_time;default:null"`
	UpdateTime *time.Time `gorm:"column:update_time;default:null"`
}

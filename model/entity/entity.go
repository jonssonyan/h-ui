package entity

import "time"

type BaseEntity struct {
	Id         *int64     `gorm:"column:id;primaryKey" json:"id"`
	CreateTime *time.Time `gorm:"column:create_time;default:null" json:"createTime"`
	UpdateTime *time.Time `gorm:"column:update_time;default:null" json:"updateTime"`
}

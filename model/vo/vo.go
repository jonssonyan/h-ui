package vo

import "time"

type BaseVo struct {
	Id         int64     `json:"id"`
	CreateTime time.Time `json:"createTime"`
}

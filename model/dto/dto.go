package dto

type BaseDto struct {
	PageNum   *int64 `json:"pageNum" form:"pageNum" validate:"required,gt=0"`      // 页号
	PageSize  *int64 `json:"pageSize" form:"pageSize" validate:"required,gt=0"`    // 页大小
	StartTime *int64 `json:"startTime" form:"startTime" validate:"omitempty,gt=0"` // 开始时间
	EndTime   *int64 `json:"endTime" form:"endTime" validate:"omitempty,gt=0"`     // 结束时间
}

type IdDto struct {
	Id *int64 `json:"id" form:"id" validate:"required,gt=0"` // 主键
}

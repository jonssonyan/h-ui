package dto

type AccountPageDto struct {
	BaseDto
	Username *string `json:"username" form:"username" validate:"omitempty,min=0,max=32,validateStr"`
	Deleted  *int64  `json:"deleted" form:"deleted" validate:"omitempty,oneof=0 1"`
}

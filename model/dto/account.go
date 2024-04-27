package dto

type AccountPageDto struct {
	BaseDto
	Username *string `json:"username" form:"username" validate:"omitempty,min=2,max=32,validateStr"`
	Deleted  *int64  `json:"deleted" form:"deleted" validate:"omitempty,oneof=0 1"`
}

type LoginDto struct {
	Username *string `json:"username" form:"username" validate:"required,min=6,max=32,validateStr"`
	Pass     *string `json:"pass" form:"pass" validate:"required,min=6,max=32,validateStr"`
}

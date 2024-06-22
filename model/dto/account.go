package dto

type AccountPageDto struct {
	BaseDto
	Username *string `json:"username" form:"username" validate:"omitempty,min=1,max=32"`
	Deleted  *int64  `json:"deleted" form:"deleted" validate:"omitempty,oneof=0 1"`
}

type LoginDto struct {
	Username *string `json:"username" form:"username" validate:"required,min=6,max=32,validateStr"`
	Pass     *string `json:"pass" form:"pass" validate:"required,min=6,max=32,validateStr"`
}

type AccountSaveDto struct {
	Username   *string `json:"username" form:"username" validate:"required,min=6,max=32,validateStr"`
	Pass       *string `json:"pass" form:"pass" validate:"required,min=6,max=32,validateStr"`
	ConPass    *string `json:"conPass" form:"conPass" validate:"required,min=6,max=32,validateStr"`
	Quota      *int64  `json:"quota" form:"quota" validate:"required,min=-1"`
	ExpireTime *int64  `json:"expireTime" form:"expireTime" validate:"required,min=0"`
	DeviceNo   *int64  `json:"deviceNo" form:"deviceNo" validate:"required,min=1"`
	Deleted    *int64  `json:"deleted" form:"deleted" validate:"required,oneof=0 1"`
}

type AccountUpdateDto struct {
	IdDto
	Username   *string `json:"username" form:"username" validate:"omitempty,min=6,max=32,validateStr"`
	Pass       *string `json:"pass" form:"pass" validate:"omitempty,min=6,max=32,validateStr"`
	ConPass    *string `json:"conPass" form:"conPass" validate:"omitempty,min=6,max=32,validateStr"`
	Quota      *int64  `json:"quota" form:"quota" validate:"omitempty,min=-1"`
	ExpireTime *int64  `json:"expireTime" form:"expireTime" validate:"omitempty,min=0"`
	DeviceNo   *int64  `json:"deviceNo" form:"deviceNo" validate:"omitempty,min=1"`
	Deleted    *int64  `json:"deleted" form:"deleted" validate:"omitempty,oneof=0 1"`
}

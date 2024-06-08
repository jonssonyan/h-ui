package dto

type Hysteria2AuthDto struct {
	Auth *string `json:"auth" form:"auth" validate:"required"`
}

type Hysteria2KickDto struct {
	Ids          []int64 `json:"ids" form:"ids" validate:"required"`
	KickUtilTime *int64  `json:"kickUtilTime" form:"kickUtilTime" validate:"required"` // 解禁时间
}

type Hysteria2VersionDto struct {
	Version *string `json:"version" form:"version" validate:"required,min=1,max=5"`
}

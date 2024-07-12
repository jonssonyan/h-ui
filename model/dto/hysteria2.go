package dto

type Hysteria2AuthDto struct {
	Addr *string `json:"addr" form:"addr" validate:"required"`
	Auth *string `json:"auth" form:"auth" validate:"required"`
	Tx   *string `json:"tx" form:"tx" validate:"required"`
}

type Hysteria2KickDto struct {
	Ids          []int64 `json:"ids" form:"ids" validate:"required"`
	KickUtilTime *int64  `json:"kickUtilTime" form:"kickUtilTime" validate:"required"` // 解禁时间
}

type Hysteria2VersionDto struct {
	Version *string `json:"version" form:"version" validate:"required,min=1,max=10"`
}

type Hysteria2SubscribeUrlDto struct {
	AccountId *int64  `json:"accountId" form:"accountId" validate:"required,gt=0"`
	Protocol  *string `json:"protocol" form:"protocol" validate:"required,min=1,max=8"`
	Host      *string `json:"host" form:"host" validate:"required,min=1,max=301"`
}

type Hysteria2UrlDto struct {
	AccountId *int64  `json:"accountId" form:"accountId" validate:"required,gt=0"`
	Hostname  *string `json:"hostname" form:"hostname" validate:"required,min=1,max=255"`
}

package dto

type Hysteria2AuthDto struct {
	Auth *string `json:"auth" form:"auth" validate:"required"`
}

type Hysteria2KickDto struct {
	Usernames []string `json:"usernames" form:"usernames" validate:"required"`
}

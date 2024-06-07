package dto

type ServerDto struct {
	Port *int64 `json:"port" form:"port" validate:"required,min=1,max=65535"`
}

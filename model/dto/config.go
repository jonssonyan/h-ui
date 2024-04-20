package dto

type ConfigDto struct {
	Keys []string `json:"keys" form:"keys" validate:"required"`
}

type ConfigUpdateDto struct {
	Key   *string `json:"username" form:"username" validate:"required,min=2,max=32"`
	Value *string `json:"value" form:"value" validate:"required,min=1,max=128"`
}

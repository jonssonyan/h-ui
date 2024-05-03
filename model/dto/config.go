package dto

type ConfigDto struct {
	Key *string `json:"key" form:"key" validate:"required,min=1,max=128"`
}

type ConfigUpdateDto struct {
	Key   *string `json:"key" form:"key" validate:"required,min=1,max=128"`
	Value *string `json:"value" form:"value" validate:"required,min=1,max=128"`
}

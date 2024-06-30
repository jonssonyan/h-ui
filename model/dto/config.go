package dto

type ConfigDto struct {
	Key *string `json:"key" form:"key" validate:"required,min=1,max=128"`
}

type ConfigsDto struct {
	Keys []string `json:"keys" form:"keys" validate:"required"`
}

type ConfigUpdateDto struct {
	Key   *string `json:"key" form:"key" validate:"required,min=1,max=128"`
	Value *string `json:"value" form:"value" validate:"required,min=0,max=128"`
}

type ConfigsUpdateDto struct {
	ConfigUpdateDtos []ConfigUpdateDto `json:"configUpdateDtos" form:"configUpdateDtos" validate:"required"`
}

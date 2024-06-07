package dto

type LogDto struct {
	NumLine *int `json:"numLine" form:"numLine" validate:"omitempty,min=1,max=300"`
}

type LogExportDto struct {
	Option *int `json:"option" form:"option" validate:"required,oneof=0 1"`
}

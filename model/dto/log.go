package dto

type LogSystemDto struct {
	NumLine *int `json:"numLine" form:"numLine" validate:"omitempty,min=1,max=300"`
}

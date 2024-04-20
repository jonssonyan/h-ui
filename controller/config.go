package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/dto"
	"h-ui/model/vo"
	"h-ui/service"
)

func UpdateConfig(c *gin.Context) {
	configUpdateDto, err := validateField(c, dto.ConfigUpdateDto{})
	if err != nil {
		return
	}
	if err = service.UpdateConfig(*configUpdateDto.Key, *configUpdateDto.Value); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func ListConfig(c *gin.Context) {
	configDto, err := validateField(c, dto.ConfigDto{})
	if err != nil {
		return
	}
	configs, err := service.ListConfig(configDto.Keys)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(configs, c)
}

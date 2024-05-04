package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/bo"
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

func UpdateConfigs(c *gin.Context) {
	configsUpdateDto, err := validateField(c, dto.ConfigsUpdateDto{})
	if err != nil {
		return
	}
	for _, item := range configsUpdateDto.ConfigUpdateDtos {
		if err = service.UpdateConfig(*item.Key, *item.Value); err != nil {
			vo.Fail(err.Error(), c)
			return
		}
	}
	vo.Success(nil, c)
}

func GetConfig(c *gin.Context) {
	configDto, err := validateField(c, dto.ConfigDto{})
	if err != nil {
		return
	}
	config, err := service.GetConfig(*configDto.Key)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	configVo := vo.ConfigVo{
		Key:   *config.Key,
		Value: *config.Value,
	}
	vo.Success(configVo, c)
}

func ListConfig(c *gin.Context) {
	configsDto, err := validateField(c, dto.ConfigsDto{})
	if err != nil {
		return
	}
	configs, err := service.ListConfig(configsDto.Keys)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	var configVos []vo.ConfigVo
	for _, item := range configs {
		configVo := vo.ConfigVo{
			Key:   *item.Key,
			Value: *item.Value,
		}
		configVos = append(configVos, configVo)
	}
	vo.Success(configVos, c)
}

func GetHysteria2Config(c *gin.Context) {
	config, err := service.GetHysteria2Config()
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(config, c)
}

func UpdateHysteria2Config(c *gin.Context) {
	hysteria2ServerConfig, err := validateField(c, bo.Hysteria2ServerConfig{})
	if err != nil {
		return
	}
	if err = service.UpdateHysteria2Config(hysteria2ServerConfig); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

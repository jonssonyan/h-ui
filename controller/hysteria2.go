package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/dto"
	"h-ui/model/vo"
	"h-ui/service"
)

func StartHysteria2(c *gin.Context) {
	if err := service.StartHysteria2(); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func StopHysteria2(c *gin.Context) {
	if err := service.StopHysteria2(); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func RestartHysteria2(c *gin.Context) {
	if err := service.RestartHysteria2(); err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

func Hysteria2Auth(c *gin.Context) {
	hysteria2AuthDto, err := validateField(c, dto.Hysteria2AuthDto{})
	if err != nil {
		return
	}
	username, err := service.Hysteria2Auth(*hysteria2AuthDto.Auth)
	if err != nil || username == "" {
		vo.Hysteria2AuthFail("", c)
		return
	}
	vo.Hysteria2AuthSuccess(username, c)
}

func Hysteria2Kick(c *gin.Context) {
	hysteria2KickDto, err := validateField(c, dto.Hysteria2KickDto{})
	if err != nil {
		return
	}
	err = service.Hysteria2Kick(hysteria2KickDto.Ids, hysteria2KickDto.KickUtilTime)
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(nil, c)
}

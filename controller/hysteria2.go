package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"h-ui/model/dto"
	"h-ui/model/vo"
	"h-ui/service"
	"h-ui/util"
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

func GetHysteria2Version(c *gin.Context) {
	version, err := util.Exec(fmt.Sprintf("%s -version", util.GetHysteria2BinPath()))
	if err != nil {
		vo.Fail(err.Error(), c)
		return
	}
	vo.Success(version, c)
}

func GetHysteria2Status(c *gin.Context) {
	running := service.Hysteria2Status()
	vo.Success(running, c)
}

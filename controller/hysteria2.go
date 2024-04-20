package controller

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/dto"
	"h-ui/model/vo"
	"h-ui/service"
)

func Hysteria2Auth(c *gin.Context) {
	hysteria2AuthDto, err := validateField(c, dto.Hysteria2AuthDto{})
	if err != nil {
		return
	}
	username, err := service.Hysteria2Auth(*hysteria2AuthDto.Auth)
	if err != nil || username != "" {
		vo.Hysteria2AuthFail("", c)
		return
	}
	vo.Hysteria2AuthSuccess(username, c)
}

package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/controller"
)

func initHysteria2AuthRouter(hysteria2Api *gin.RouterGroup) {
	hysteria2 := hysteria2Api.Group("/hysteria2")
	{
		hysteria2.POST("/auth", controller.Hysteria2Auth)
	}
}

func initHysteria2Router(hysteria2Api *gin.RouterGroup) {
	hysteria2 := hysteria2Api.Group("/hysteria2")
	{
		hysteria2.POST("/kick", controller.Hysteria2Kick)
	}
}

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
		hysteria2.POST("/startHysteria2", controller.StartHysteria2)
		hysteria2.POST("/stopHysteria2", controller.StopHysteria2)
		hysteria2.POST("/restartHysteria2", controller.RestartHysteria2)
		hysteria2.POST("/countOnline", controller.CountOnline)
		hysteria2.POST("/kick", controller.Hysteria2Kick)
	}
}

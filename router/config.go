package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/controller"
)

func initConfigRouter(configApi *gin.RouterGroup) {
	config := configApi.Group("/config")
	{
		config.POST("/updateConfig", controller.UpdateConfig)
		config.GET("/getConfig", controller.GetConfig)
		config.GET("/getHysteria2Config", controller.GetHysteria2Config)
		config.POST("/updateHysteria2Config", controller.UpdateHysteria2Config)
	}
}

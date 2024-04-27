package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/controller"
)

func initConfigRouter(configApi *gin.RouterGroup) {
	config := configApi.Group("/config")
	{
		config.POST("/updateConfig", controller.UpdateConfig)
		config.GET("/listConfig", controller.ListConfig)
		config.POST("/getHysteria2Config", controller.GetHysteria2Config)
	}
}

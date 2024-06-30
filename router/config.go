package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/controller"
)

func initConfigRouter(configApi *gin.RouterGroup) {
	config := configApi.Group("/config")
	{
		config.POST("/updateConfigs", controller.UpdateConfigs)
		config.GET("/getConfig", controller.GetConfig)
		config.POST("/listConfig", controller.ListConfig)
		config.GET("/getHysteria2Config", controller.GetHysteria2Config)
		config.POST("/updateHysteria2Config", controller.UpdateHysteria2Config)
		config.POST("/exportHysteria2Config", controller.ExportHysteria2Config)
		config.POST("/importHysteria2Config", controller.ImportHysteria2Config)
		config.POST("/exportConfig", controller.ExportConfig)
		config.POST("/importConfig", controller.ImportConfig)
		config.GET("/hysteria2AcmePath", controller.Hysteria2AcmePath)
		config.POST("/restartServer", controller.RestartServer)
	}
}

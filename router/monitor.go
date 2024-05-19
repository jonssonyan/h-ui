package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/controller"
)

func initMonitorRouter(accountApi *gin.RouterGroup) {
	account := accountApi.Group("/monitor")
	{
		account.GET("/system", controller.MonitorSystem)
	}
}

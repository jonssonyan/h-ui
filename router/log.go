package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/controller"
)

func initLogRouter(accountApi *gin.RouterGroup) {
	account := accountApi.Group("/log")
	{
		account.GET("/system", controller.LogSystem)
		account.GET("/hysteria2", controller.LogHysteria2)
	}
}

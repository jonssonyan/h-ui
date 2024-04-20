package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/middleware"
)

func Router(router *gin.Engine) {
	router.Use(middleware.LogHandler())
	api := router.Group("/hui")
	{
		initAccountRouter(api)
		initHysteria2Router(api)
	}
}

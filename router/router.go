package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/frontend"
	"h-ui/middleware"
)

func Router(router *gin.Engine) {
	router.Use(middleware.LogHandler())

	frontend.InitFrontend(router)

	api := router.Group("/hui")
	{
		initAccountRouter(api)
		initConfigRouter(api)
		initHysteria2Router(api)
	}
}

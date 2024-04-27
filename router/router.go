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
		initAuthRouter(api)
		initHysteria2AuthRouter(api)
	}

	router.Use(middleware.JWTHandler())

	huiApi := router.Group("/hui")
	{
		initAccountRouter(huiApi)
		initConfigRouter(huiApi)
		initHysteria2Router(api)
	}
}

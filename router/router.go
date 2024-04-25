package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/frontend"
	"h-ui/middleware"
)

func Router(router *gin.Engine) {
	router.Use(middleware.LogHandler())

	frontend.InitFrontend(router)

	authApi := router.Group("/hui")
	{
		initAuthRouter(authApi)
	}

	router.Use(middleware.JWTHandler())

	huiApi := router.Group("/hui")
	{
		initAccountRouter(huiApi)
		initConfigRouter(huiApi)
		initHysteria2Router(huiApi)
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/frontend"
	"h-ui/middleware"
)

func Router(router *gin.Engine) {

	router.Use(middleware.LogHandler(), middleware.RateLimiterHandler())

	frontend.InitFrontend(router)

	authApi := router.Group("/hui")
	{
		initAuthRouter(authApi)
		initHysteria2AuthRouter(authApi)
	}

	router.Use(middleware.JWTHandler())

	router.Use(middleware.AdminHandler())

	huiAdminApi := router.Group("/hui")
	{
		initAccountAdminRouter(huiAdminApi)
		initConfigRouter(huiAdminApi)
		initHysteria2Router(huiAdminApi)
		initLogRouter(huiAdminApi)
		initMonitorRouter(huiAdminApi)
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/frontend"
	"h-ui/middleware"
)

func Router(router *gin.Engine) {
	// 日志和单机限流
	router.Use(middleware.LogHandler(), middleware.RateLimiterHandler())

	frontend.InitFrontend(router)

	// 不需要鉴权
	authApi := router.Group("/hui")
	{
		initAuthRouter(authApi)
		initHysteria2AuthRouter(authApi)
	}

	// 认证
	router.Use(middleware.JWTHandler())

	// 鉴权
	router.Use(middleware.AdminHandler())

	huiAdminApi := router.Group("/hui")
	{
		initAccountAdminRouter(huiAdminApi)
		initConfigRouter(huiAdminApi)
		initHysteria2Router(huiAdminApi)
		initLogRouter(huiAdminApi)
	}
}

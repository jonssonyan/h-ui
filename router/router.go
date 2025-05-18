package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/frontend"
	"h-ui/middleware"
	"strings"
)

func Router(router *gin.Engine, huiWebContext *string) {
	// global context
	relativePath := "/"
	if huiWebContext != nil && strings.HasPrefix(*huiWebContext, "/") {
		relativePath = *huiWebContext
	}
	globalGroup := router.Group(relativePath)
	{
		globalGroup.Use(middleware.FilterHandler(), middleware.LogHandler(), middleware.RateLimiterHandler())

		frontend.InitFrontend(router, relativePath)

		authApi := globalGroup.Group("/hui")
		{
			initAuthRouter(authApi)
			initHysteria2AuthRouter(authApi)
		}

		globalGroup.Use(middleware.JWTHandler())

		globalGroup.Use(middleware.AdminHandler())

		huiAdminApi := globalGroup.Group("/hui")
		{
			initAccountAdminRouter(huiAdminApi)
			initConfigRouter(huiAdminApi)
			initHysteria2Router(huiAdminApi)
			initLogRouter(huiAdminApi)
			initMonitorRouter(huiAdminApi)
		}
	}
}

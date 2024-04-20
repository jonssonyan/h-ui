package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/controller"
)

func initAccountRouter(accountApi *gin.RouterGroup) {
	account := accountApi.Group("/account")
	{
		account.GET("/pageAccount", controller.PageAccount)
	}
}

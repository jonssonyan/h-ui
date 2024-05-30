package router

import (
	"github.com/gin-gonic/gin"
	"h-ui/controller"
)

func initAccountAdminRouter(accountApi *gin.RouterGroup) {
	account := accountApi.Group("/account")
	{
		account.GET("/pageAccount", controller.PageAccount)
		account.POST("/saveAccount", controller.SaveAccount)
		account.POST("/deleteAccount", controller.DeleteAccount)
		account.POST("/updateAccount", controller.UpdateAccount)
		account.GET("/getAccountInfo", controller.GetAccountInfo)
		account.GET("/getAccount", controller.GetAccount)
		account.POST("/exportAccount", controller.ExportAccount)
		account.POST("/releaseKickAccount", controller.ReleaseKickAccount)
		account.POST("/importAccount", controller.ImportAccount)
	}
}

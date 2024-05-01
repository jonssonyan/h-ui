package middleware

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/constant"
	"h-ui/model/vo"
	"h-ui/service"
	"h-ui/util"
)

func AdminHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		myClaims, err := service.ParseToken(service.GetToken(c))
		if err != nil {
			vo.Fail(err.Error(), c)
			c.Abort()
			return
		}
		if !util.ArrContain(myClaims.AccountBo.Roles, "admin") {
			vo.Fail(constant.ForbiddenError, c)
			c.Abort()
			return
		}
		c.Next()
	}
}

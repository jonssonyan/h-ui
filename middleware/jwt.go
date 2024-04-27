package middleware

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/constant"
	"h-ui/model/vo"
	"h-ui/service"
	"strings"
)

func JWTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			vo.Fail(constant.UnauthorizedError, c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			vo.Fail(constant.IllegalTokenError, c)
			c.Abort()
			return
		}
		myClaims, err := service.ParseToken(parts[1])
		if err != nil {
			vo.Fail(err.Error(), c)
			c.Abort()
			return
		}
		if myClaims.AccountBo.Deleted != 0 {
			vo.Fail("this account has been disabled", c)
			c.Abort()
			return
		}
		c.Next()
	}
}

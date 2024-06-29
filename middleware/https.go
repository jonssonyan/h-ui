package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"h-ui/model/constant"
	"h-ui/model/vo"
	"h-ui/service"
	"net/http"
	"strings"
)

func RedirectHttpsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		server, err := service.NewServer()
		if err != nil {
			vo.Fail(err.Error(), c)
			c.Abort()
			return
		}
		if server.IsHttps() && c.Request.TLS == nil {
			httpsPortConfig, err := service.GetConfig(constant.HUIWebHttpsPort)
			if err != nil {
				vo.Fail(err.Error(), c)
				c.Abort()
				return
			}
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://%s:%s%s", strings.Split(c.Request.Host, ":")[0], *httpsPortConfig.Value, c.Request.RequestURI))
			c.Abort()
			return
		}
		c.Next()
	}
}

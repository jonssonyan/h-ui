package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
			server, err := service.NewServer()
			if err != nil {
				vo.Fail(err.Error(), c)
				c.Abort()
				return
			}
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://%s:%d%s", strings.Split(c.Request.Host, ":")[0], server.HttpsPort, c.Request.RequestURI))
			c.Abort()
			return
		}
		c.Next()
	}
}

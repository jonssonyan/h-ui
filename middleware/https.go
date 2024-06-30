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
		httpsPort := server.GetHttps()
		if httpsPort > 0 && c.Request.TLS == nil {
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://%s:%d%s", strings.Split(c.Request.Host, ":")[0], httpsPort, c.Request.RequestURI))
			c.Abort()
			return
		}
		c.Next()
	}
}

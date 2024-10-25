package middleware

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/vo"
	"net/http"
	"regexp"
)

func FilterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		matched, err := regexp.MatchString(`(?i)fofa|shodan|curl|wget`, c.Request.UserAgent())
		if err != nil {
			vo.Fail("Internal error", c)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if matched {
			vo.Fail("Forbidden: Scanning tools are not allowed", c)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

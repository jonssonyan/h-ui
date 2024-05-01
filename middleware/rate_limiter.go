package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"h-ui/model/vo"
)

var limit *limiter.Limiter

func init() {
	limit = tollbooth.NewLimiter(5, nil)
}

func RateLimiterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(limit, c.Writer, c.Request)
		if httpError != nil {
			vo.Fail("click too fast", c)
			c.Abort()
			return
		}
		c.Next()
	}
}

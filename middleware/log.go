package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"h-ui/model/constant"
	"time"
)

func InitLog() {
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   constant.SystemLogPath,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	})
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logrus.SetLevel(logrus.WarnLevel)
}

func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		endTime := time.Now()
		statusCode := c.Writer.Status()
		latencyTime := endTime.Sub(startTime)
		clientIP := c.ClientIP()
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI

		logrus.WithFields(logrus.Fields{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
		}).Info()
		c.Next()
	}
}

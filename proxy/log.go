package proxy

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerPlus struct {
	Logger *logrus.Logger
}

func NewLoggerPlus(filename string, maxSize int, maxBackups int, maxAge int, compress bool) *LoggerPlus {
	var logger *logrus.Logger
	logger.SetOutput(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
		LocalTime:  true,
	})
	logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logger.SetLevel(logrus.WarnLevel)
	return &LoggerPlus{
		Logger: logger,
	}
}

func (loggerPlus *LoggerPlus) Infof(format string, args ...interface{}) {
	loggerPlus.Logger.Infof(format, args)
}

func (loggerPlus *LoggerPlus) Errorf(format string, args ...interface{}) {
	loggerPlus.Logger.Errorf(format, args)
}

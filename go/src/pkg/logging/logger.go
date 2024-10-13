package logging

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func Initialize() {
	if logger != nil {
		return
	}
	ginMode := gin.Mode()
	var (
		l *zap.Logger
	)
	switch ginMode {
	case gin.DebugMode:
		l, _ = zap.NewDevelopment()
	case gin.TestMode, gin.ReleaseMode:
		l, _ = zap.NewProduction()
	default:
		l, _ = zap.NewDevelopment()
	}
	logger = l.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		Initialize()
	}
	return logger
}

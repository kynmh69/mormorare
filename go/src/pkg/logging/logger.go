package logging

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	logger    *zap.SugaredLogger
	zapLogger *zap.Logger
)

func Initialize() {
	if logger != nil {
		return
	}
	ginMode := gin.Mode()
	switch ginMode {
	case gin.DebugMode:
		zapLogger, _ = zap.NewDevelopment()
	case gin.TestMode, gin.ReleaseMode:
		zapLogger, _ = zap.NewProduction()
	default:
		zapLogger, _ = zap.NewDevelopment()
	}
	logger = zapLogger.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	if logger == nil {
		Initialize()
	}
	return logger
}

func GetZapLogger() *zap.Logger {
	if zapLogger == nil {
		Initialize()
	}
	return zapLogger
}

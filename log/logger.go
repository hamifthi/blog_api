package log

import (
	"go.uber.org/zap"
)

type Logger interface {
	Info(msg string, values ...any)
	Warn(msg string, values ...any)
	Error(msg string, values ...any)
	Debug(msg string, values ...any)
}

func ProductionConfig() zap.Config {
	return zap.NewProductionConfig()
}

func DevelopmentConfig() zap.Config {
	return zap.NewDevelopmentConfig()
}

package utils

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger() {

	env := config.AppConfig.AppEnv
	var err error

	if env == "development" {
		cfg := zap.NewDevelopmentConfig()
		cfg.DisableStacktrace = true
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = cfg.Build()
	} else {
		cfg := zap.NewProductionConfig()
		cfg.DisableStacktrace = false
		logger, err = cfg.Build()
	}

	if err != nil {
		panic("failed to initialize logger")
	}
}
func GetLogger() *zap.Logger {
	return logger
}

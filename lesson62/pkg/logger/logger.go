package logger

import (
	"go.uber.org/zap"
)

func DevelopmentConfig(file string) zap.Config {
	configZap := zap.NewDevelopmentConfig()

	configZap.OutputPaths = []string{"stdout", file}
	configZap.ErrorOutputPaths = []string{"strerr"}

	return configZap
}

func New(level, fileName string) (*zap.Logger, error) {

	configZap := DevelopmentConfig(fileName)

	switch level {
	case "debug":
		configZap.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		configZap.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		configZap.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		configZap.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		configZap.Level = zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		configZap.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		configZap.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		configZap.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	return configZap.Build()
}

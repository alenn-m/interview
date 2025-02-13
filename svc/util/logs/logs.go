package logs

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func GetLogger() Logger {
	return newLogger()
}

// newLogger sets up logger
func newLogger() Logger {
	env := os.Getenv("ENV")

	zapConfig := zap.NewProductionConfig()
	if env == "development" {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapConfig.OutputPaths = append(zapConfig.OutputPaths, fmt.Sprintf("%s/interview.log", os.Getenv("HOME")))

	zapLogger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	logger := newSugaredLogger(zapLogger)

	return *logger
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

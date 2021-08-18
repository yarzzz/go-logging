package logging

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewConsoleLogger(name string) *Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zap.CombineWriteSyncers(
				zapcore.AddSync(os.Stdout),
			),
			zapcore.InfoLevel,
		),
		zap.AddCaller(),
	)
	logger = logger.Named(name)
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

func NewFileLogger(name string, path string) *Logger {
	f, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zap.CombineWriteSyncers(
				zapcore.AddSync(f),
			),
			zapcore.InfoLevel,
		),
		zap.AddCaller(),
	)
	logger = logger.Named(name)
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

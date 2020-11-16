package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var singleton = Factory{
	logger: func() *zap.Logger {
		logger, _ := zap.NewProduction()
		return logger
	}(),
}

func Info(msg string, fields ...zapcore.Field) {
	singleton.Bg().Info(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	singleton.Bg().Error(msg, fields...)

}

func Fatal(msg string, fields ...zapcore.Field) {
	singleton.Bg().Fatal(msg, fields...)
}

func For(ctx context.Context) Logger {
	return singleton.For(ctx)
}

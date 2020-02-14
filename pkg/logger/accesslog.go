package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var accesslog *zap.Logger

var AccessLogEncoderConfig = zapcore.EncoderConfig{
	TimeKey:       "",
	LevelKey:      "",
	NameKey:       "",
	CallerKey:     "",
	StacktraceKey: "",
	EncodeTime:    zapcore.ISO8601TimeEncoder,
}

func initAccessLog() {
	accesslogHook := lumberjack.Logger{
		Filename:   "logs/access.log",
		MaxSize:    128, //MB
		MaxBackups: 10,
		MaxAge:     45,
		Compress:   true,
	}

	accesslog = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(AccessLogEncoderConfig),
		zapcore.AddSync(&accesslogHook),
		zap.InfoLevel,
	))
}

func AccessLog(fields ...zap.Field) {
	accesslog.Info("", fields...)
}

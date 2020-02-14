package logger

import (
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger   *zap.Logger
	logLevel = zap.NewAtomicLevel()
)

var DefaultEncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "ts",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func InitLogger(filename, level string) {
	hook := lumberjack.Logger{
		Filename:   filename,
		MaxSize:    128, //MB
		MaxBackups: 10,
		MaxAge:     45,
		Compress:   true,
	}

	switch strings.ToLower(level) {
	case "debug":
		logLevel.SetLevel(zapcore.DebugLevel)
	case "info":
		logLevel.SetLevel(zapcore.InfoLevel)
	case "warn":
		fallthrough
	case "warning":
		logLevel.SetLevel(zapcore.WarnLevel)
	case "error":
		logLevel.SetLevel(zapcore.ErrorLevel)
	}

	logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(DefaultEncoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		logLevel,
	), zap.AddCaller())

	initAccessLog()

	logger.Info("Zap Logger init success")
}

func GetLevel() zap.AtomicLevel {
	return logLevel
}

func LG() *zap.Logger {
	return logger
}

func Sugar() *zap.SugaredLogger {
	return logger.Sugar()
}

func Close() {
	if logger != nil {
		logger.Sync()
	}
	if accesslog != nil {
		accesslog.Sync()
	}
}

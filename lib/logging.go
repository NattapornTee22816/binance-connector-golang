package lib

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type LogLevel int

var (
	LogLevelInfo  = LogLevel(1)
	LogLevelDebug = LogLevel(2)
	LogLevelTrace = LogLevel(3)
)

type BinanceLogger struct {
	*zap.Logger
	*BinanceLoggerHelper
}

type BinanceLoggerHelper struct {
	moduleName string
	Level      LogLevel
}

func NewLogger(moduleName string, logLevel LogLevel) *BinanceLogger {
	helper := &BinanceLoggerHelper{
		moduleName: moduleName,
		Level:      logLevel,
	}

	logger, _ := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    helper.capitalLevelEncoder,
			EncodeTime:     helper.timeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		DisableCaller:    true,
	}.Build()

	return &BinanceLogger{
		logger,
		helper,
	}
}

func (r *BinanceLoggerHelper) timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s] %s", r.moduleName, t.Format("2006-01-02 15:04:05")))
}

func (r *BinanceLoggerHelper) capitalLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", l.CapitalString()[:4]))
}

func (r *BinanceLoggerHelper) CanDebug() bool {
	return r.Level >= LogLevelDebug
}

func (r *BinanceLogger) CanTrace() bool {
	return r.Level >= LogLevelTrace
}

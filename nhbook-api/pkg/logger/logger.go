package logger

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	App     *zap.Logger
	Request *zap.Logger
	Error   *zap.Logger
}

func parseLevel(levelStr string) (zap.AtomicLevel, error) {
	var level zap.AtomicLevel
	switch levelStr {
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		level = zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		return level, errors.New("level invalid")
	}
	return level, nil
}

func InitLogger(cnf *ConfigLogger) (*Logger, error) {
	// Parser Level
	level, err := parseLevel(cnf.Level)
	if err != nil {
		return nil, err
	}

	// Encoder Config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// App core logger
	appCore := zapcore.NewCore(encoder, GetWrite(cnf.AppLog, cnf), level)

	// Request core logger
	requestCore := zapcore.NewCore(encoder, GetWrite(cnf.RequestLog, cnf), level)

	// Error core logger
	errorCore := zapcore.NewCore(encoder, GetWrite(cnf.ErrorLog, cnf), zapcore.ErrorLevel)

	return &Logger{
		App:     zap.New(appCore, zap.AddCaller()),
		Request: zap.New(requestCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
		Error:   zap.New(errorCore, zap.AddCaller(), zap.AddStacktrace(level)),
	}, nil
}

// Sync
func Sync(logger *Logger) {
	_ = logger.App.Sync()
	_ = logger.Request.Sync()
	_ = logger.Error.Sync()
}

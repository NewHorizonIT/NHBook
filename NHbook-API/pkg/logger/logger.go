package logger

import (
	"os"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	// Config of Encoder
	c := global.Config
	encodeConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		CallerKey:    "caller",
		MessageKey:   "msg",
		LevelKey:     "level",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
	}
	// Create JSON Encoder
	encodeLogger := zapcore.NewJSONEncoder(encodeConfig)

	// Setup level log
	var level zapcore.Level
	switch c.Env {
	case "dev":
		level = zapcore.DebugLevel
	case "pro":
		level = zapcore.InfoLevel
	default:
		level = zapcore.InfoLevel
	}
	// Setup type write log
	writer := zapcore.NewMultiWriteSyncer(writeSync(), zapcore.AddSync(os.Stdout))
	// Create Core logger
	core := zapcore.NewCore(encodeLogger, writer, level)
	// Create Logger
	logger := zap.New(core).WithOptions(zap.AddCaller())

	global.Logger = logger

}

// Setup Lumberjack
func writeSync() zapcore.WriteSyncer {
	lc := global.Config.Logger
	lumberjackLogger := &lumberjack.Logger{
		Filename:   lc.FileName,
		MaxSize:    lc.MaxSize,
		MaxBackups: lc.MaxBackups,
		MaxAge:     lc.MaxAge,
		Compress:   lc.Compress,
	}
	return zapcore.AddSync(lumberjackLogger)
}

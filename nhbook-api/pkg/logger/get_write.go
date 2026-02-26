package logger

import (
	"os"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func GetWrite(file string, cnf *ConfigLogger) zapcore.WriteSyncer {
	// Lumberjack config
	lumberjackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    cnf.MaxSize,
		MaxBackups: cnf.MaxBackups,
		MaxAge:     cnf.MaxAge,
		Compress:   cnf.Compress,
	}

	fileWriteSyncer := zapcore.AddSync(lumberjackLogger)
	consoleWriteSyncer := zapcore.AddSync(os.Stdout)

	var multipleWriteSyncer zapcore.WriteSyncer
	if cnf.EnableConsole {
		multipleWriteSyncer = zapcore.NewMultiWriteSyncer(fileWriteSyncer, consoleWriteSyncer)
	} else {
		multipleWriteSyncer = zapcore.NewMultiWriteSyncer(fileWriteSyncer)
	}

	return multipleWriteSyncer
}

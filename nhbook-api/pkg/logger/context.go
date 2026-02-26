package logger

import (
	"context"

	"go.uber.org/zap"
)

type ctxKey struct {
}

func WithContext(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, log)
}

func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return logger
	}

	return zap.L()
}

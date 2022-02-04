package ctxzap

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxzap struct{}

func WithLogger(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxzap{}, l)
}

func Logger(ctx context.Context) *zap.Logger {
	v := logger(ctx)
	if v == nil {
		v = zap.NewNop()
	}
	return v
}

func logger(ctx context.Context) *zap.Logger {
	v, _ := ctx.Value(ctxzap{}).(*zap.Logger)
	return v
}

func With(ctx context.Context, fields ...zapcore.Field) context.Context {
	return WithLogger(ctx, Logger(ctx).With(fields...))
}

func WithOptions(ctx context.Context, options ...zap.Option) context.Context {
	return WithLogger(ctx, Logger(ctx).WithOptions(options...))
}

func Named(ctx context.Context, name string) context.Context {
	return WithLogger(ctx, Logger(ctx).Named(name))
}

func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	if l := logger(ctx); l != nil {
		l.Debug(msg, fields...)
	}
}

func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	if l := logger(ctx); l != nil {
		l.Info(msg, fields...)
	}
}

func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	if l := logger(ctx); l != nil {
		l.Warn(msg, fields...)
	}
}

func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	if l := logger(ctx); l != nil {
		l.Error(msg, fields...)
	}
}

func DPanic(ctx context.Context, msg string, fields ...zapcore.Field) {
	Logger(ctx).DPanic(msg, fields...)
}

func Panic(ctx context.Context, msg string, fields ...zapcore.Field) {
	Logger(ctx).Panic(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zapcore.Field) {
	Logger(ctx).Fatal(msg, fields...)
}

func Check(ctx context.Context, lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	if l := logger(ctx); l != nil {
		return l.Check(lvl, msg)
	}
	return nil
}

func Sync(ctx context.Context) error {
	if l := logger(ctx); l != nil {
		return l.Sync()
	}
	return nil
}

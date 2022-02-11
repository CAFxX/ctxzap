package ctxzap

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxzap struct{}

func WithLogger(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxzap{}, l.WithOptions(zap.AddCallerSkip(1)))
}

func Logger(ctx context.Context) *zap.Logger {
	return logger(ctx, true).WithOptions(zap.AddCallerSkip(-1))
}

func logger(ctx context.Context, force bool) *zap.Logger {
	v, _ := ctx.Value(ctxzap{}).(*zap.Logger)
	if v == nil && force {
		v = zap.NewNop().WithOptions(zap.AddCallerSkip(1))
	}
	return v
}

func With(ctx context.Context, fields ...zapcore.Field) context.Context {
	if l := logger(ctx, false); l != nil {
		return WithLogger(ctx, l.With(fields...))
	}
	return ctx
}

func WithOptions(ctx context.Context, options ...zap.Option) context.Context {
	return WithLogger(ctx, Logger(ctx).WithOptions(options...))
}

func Named(ctx context.Context, name string) context.Context {
	return WithLogger(ctx, Logger(ctx).Named(name))
}

func Debug(ctx context.Context, msg string, fields ...zapcore.Field) {
	if l := logger(ctx, false); l != nil {
		l.Debug(msg, fields...)
	}
}

func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	if l := logger(ctx, false); l != nil {
		l.Info(msg, fields...)
	}
}

func Warn(ctx context.Context, msg string, fields ...zapcore.Field) {
	if l := logger(ctx, false); l != nil {
		l.Warn(msg, fields...)
	}
}

func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	if l := logger(ctx, false); l != nil {
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
	if l := logger(ctx, lvl >= zap.PanicLevel); l != nil {
		return l.Check(lvl, msg)
	}
	return nil
}

func Sync(ctx context.Context) error {
	if l := logger(ctx, false); l != nil {
		return l.Sync()
	}
	return nil
}

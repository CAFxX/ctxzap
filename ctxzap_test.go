package ctxzap

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestCtxzap(t *testing.T) {
	ctx := context.Background()
	Info(ctx, "info")
}

func TestCtxzapWith(t *testing.T) {
	ctx := context.Background()
	ctx = With(ctx, zap.String("foo", "bar"))
	Info(ctx, "info")
}

func TestCtxzapCheck(t *testing.T) {
	ctx := context.Background()
	if ce := Check(ctx, zapcore.ErrorLevel, "msg"); ce != nil {
		ce.Write()
	}
}

func TestCtxzapPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("no panic")
		}
	}()
	ctx := context.Background()
	Panic(ctx, "ohnoes")
}

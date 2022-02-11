package ctxzap

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ExampleInfo() {
	ctx := context.Background()
	Info(ctx, "not logged")

	l := zap.NewExample()
	ctx = WithLogger(ctx, l)
	Info(ctx, "logged")

	// Output:
	// {"level":"info","msg":"logged"}
}

func ExampleWith() {
	ctx := context.Background()
	ctx = With(ctx, zap.String("foo", "bar"))
	Info(ctx, "not logged")

	ctx = WithOptions(ctx, zap.WrapCore(func(zapcore.Core) zapcore.Core {
		return zap.NewExample().Core()
	}))
	Info(ctx, "logged")

	// Output:
	// {"level":"info","msg":"logged"}
}

func TestCheck(t *testing.T) {
	ctx := context.Background()
	if ce := Check(ctx, zapcore.ErrorLevel, "msg"); ce != nil {
		t.Fatal("Check returned non-nil")
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("no panic")
		}
	}()
	ctx := context.Background()
	Panic(ctx, "ohnoes")
}

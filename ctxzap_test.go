package ctxzap

import (
	"context"
	"testing"

	"go.uber.org/zap"
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

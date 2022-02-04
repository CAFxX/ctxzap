# ctxzap

Attach your zap.Logger to context.Context. [![Package documentation](https://pkg.go.dev/badge/github.com/CAFxX/ctxzap)](https://pkg.go.dev/github.com/CAFxX/ctxzap)

## Usage

```
// If ctx contains a zap.Logger, this is equivalent to a call to Info on the zap.Logger,
// if the ctx does not contain a zap.Logger, this is equivalent to a call to Info on a Nop zap.Logger.
ctxzap.Info(ctx, "message", zap.String("key", "val"))

// It is possible to obtain a derived context with the specified zapcore.Fields.
// This is equivalent to calling With() on the zap.Logger in the context, and creating a new context
// with the zap.Logger returned by With()
ctx = ctxzap.With(ctx, zap.String("foo", "bar"))

// It is possible to extract the zap.Logger in the context with Logger, and attach a custom logger
// to a context with WithLogger.
logger := ctxzap.Logger(ctx)

// ...modify logger as needed...

ctx = ctxzap.WithLogger(ctx, logger)
```

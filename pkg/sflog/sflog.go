package sflog

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"io"
	"log"
	"os"
)

type ctxLogger struct{}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l hclog.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) hclog.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(hclog.Logger); ok {
		return l
	}
	return hclog.Default()
}

func LoggerWithWriter(loggerName string, out io.Writer) hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name:            loggerName,
		Output:          out,
		JSONFormat:      false,
		IncludeLocation: false,
		TimeFormat:      "2006-01-02 15:04:05",
		Color:           hclog.ColorOff,
		Level:           hclog.Debug,
	})
}

func LoggerFromEnv(loggerName string) hclog.Logger {
	return LoggerWithWriter(loggerName, os.Stdout)
}

func SetDefaultLogger(logger hclog.Logger) {
	hclog.SetDefault(logger)

	log.SetOutput(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))
	log.SetPrefix("")
	log.SetFlags(0)
}

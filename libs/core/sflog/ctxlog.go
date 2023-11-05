package sflog

import (
	"context"
	"github.com/selflow/selflow/libs/selflow-daemon/sfenvironment"
	"io"
	"log/slog"
	"os"
)

var (
	defaultLoggerHandler = loggerHandlerFromEnv(os.Stdout)
)

type loggerContextKey struct{}

var key = loggerContextKey{}

func GetLogLevelFromString(level string) slog.Leveler {
	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "ERROR":
		return slog.LevelError
	case "WARN":
		return slog.LevelWarn
	}
	return slog.LevelInfo
}

func loggerHandlerFromEnv(writer io.Writer) slog.Handler {
	handlerOptions := &slog.HandlerOptions{
		Level: GetLogLevelFromString(sfenvironment.LogLevel),
	}

	if sfenvironment.UseJsonLogs {
		return slog.NewJSONHandler(writer, handlerOptions)
	}

	// Use normal logger by default for debug config
	return slog.NewTextHandler(writer, handlerOptions)
}

type ContextLogHandler struct{}

func NewLoggerWithWriter(writer io.Writer) *slog.Logger {
	return slog.New(loggerHandlerFromEnv(writer))
}

func (l *ContextLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	handler := getLogHandlerFromContext(ctx)
	return handler.Enabled(ctx, level)
}

// WithAttrs exists to validate implementation of `slog.Handler`
//
// WARNING: Shouldn't be used
func (l *ContextLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return defaultLoggerHandler.WithAttrs(attrs)
}

// WithGroup exists to validate implementation of `slog.Handler`
//
// WARNING: Shouldn't be used
func (l *ContextLogHandler) WithGroup(name string) slog.Handler {
	//WARN: shouldn't be used
	return defaultLoggerHandler.WithGroup(name)
}

func getLogHandlerFromContext(ctx context.Context) slog.Handler {
	if logger, ok := ctx.Value(key).(slog.Handler); ok {
		return logger
	}

	return defaultLoggerHandler
}

func GetLoggerFromContext(ctx context.Context) *slog.Logger {
	return slog.New(getLogHandlerFromContext(ctx))
}

func (l *ContextLogHandler) Handle(ctx context.Context, record slog.Record) error {

	handler := getLogHandlerFromContext(ctx)
	return handler.Handle(ctx, record)
}

func AddArgsToContextLogger(ctx context.Context, attrs ...slog.Attr) context.Context {
	handler := getLogHandlerFromContext(ctx)
	return ResetContextLogHandler(ctx, handler, attrs...)
}

func Init(attrs ...slog.Attr) {
	defaultLoggerHandler = defaultLoggerHandler.WithAttrs(attrs)
	logger := slog.New(&ContextLogHandler{})

	slog.SetDefault(logger)
}

func ResetContextLogHandler(ctx context.Context, handler slog.Handler, attrs ...slog.Attr) context.Context {
	handler = handler.WithAttrs(attrs)
	return context.WithValue(ctx, key, handler)
}

func WriterFromLogger(logger *slog.Logger, level slog.Level, args ...any) io.Writer {
	return LoggerWriter{
		logger:        logger,
		inferredLevel: level,
		args:          args,
	}
}

type LoggerWriter struct {
	logger        *slog.Logger
	inferredLevel slog.Level
	args          []any
}

func (l LoggerWriter) Write(p []byte) (n int, err error) {
	l.logger.Log(context.Background(), l.inferredLevel, string(p), l.args...)
	return len(p), err
}

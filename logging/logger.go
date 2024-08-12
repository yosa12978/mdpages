package logging

import (
	"io"
	"log/slog"
)

type Logger interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
}

func NewLogger(w io.Writer) Logger {
	return &slogLogger{
		logger: slog.New(slog.NewJSONHandler(w, nil)),
	}
}

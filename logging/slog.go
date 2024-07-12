package logging

import "log/slog"

type slogLogger struct {
	logger *slog.Logger
}

func (l *slogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *slogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *slogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *slogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

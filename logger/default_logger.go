package logger

import (
	"fmt"
	"log"
	"log/slog"
)

type DefaultLogger struct {
	logger *slog.Logger
}

func NewDefaultLogger() Logger {
	return DefaultLogger{logger: slog.Default()}
}

// Implementing all methods of logger.Logger to DefaultLogger
func (l DefaultLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debug(msg, keysAndValues...)
}

func (l DefaultLogger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, keysAndValues...)
}

func (l DefaultLogger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warn(msg, keysAndValues...)
}

func (l DefaultLogger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, keysAndValues...)
}

func (l DefaultLogger) With(args ...interface{}) Logger {
	return DefaultLogger{logger: l.logger.With(args...)}
}

func (l DefaultLogger) Sync() error {
	return nil
}

func (l DefaultLogger) Panic(args ...interface{}) {
	l.logger.Error("PANIC", args...)
}

func (l DefaultLogger) Fatalf(msg string, keysAndValues ...interface{}) {
	log.Fatal(msg, keysAndValues)
}

func (l DefaultLogger) Fatal(msg string, keysAndValues ...interface{}) {
	log.Fatal(msg, keysAndValues)
}

func (l DefaultLogger) Infof(template string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(template, args...))
}
func (l DefaultLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(template, args...))
}
func (l DefaultLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(template, args...))
}
func (l DefaultLogger) Panicf(template string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(template, args...))
}
func (l DefaultLogger) Warnf(template string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf(template, args...))
}

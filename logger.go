package mslog

import (
	"errors"
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	config *Config
	logger *logrus.Logger
	sentry *sentry.Hub
}

func NewLogger(config *Config, lr *logrus.Logger, sentry *sentry.Hub) *Logger {
	logger := &Logger{
		config: config,
		logger: lr,
		sentry: sentry,
	}

	return logger
}

func (logger *Logger) Info(args ...interface{}) {
	if logger.CanLog() {
		logger.logger.Info(args...)
	}
}

func (logger *Logger) Debug(args ...interface{}) {
	if logger.CanLog() {
		logger.logger.Debug(args...)
	}
}

func (logger *Logger) Trace(args ...interface{}) {
	if logger.CanLog() {
		logger.logger.Trace(args...)
	}
}

func (logger *Logger) Warn(args ...interface{}) {
	if logger.CanSentry() {
		logger.SentryWarn(args...)
	}

	if logger.CanLog() {
		logger.logger.Warn(args...)
	}
}

func (logger *Logger) Error(args ...interface{}) {
	if logger.CanSentry() {
		logger.SentryException(errors.New(fmt.Sprint(args...)))
	}

	if logger.CanLog() {
		logger.logger.Error(args...)
	}
}

func (logger *Logger) Fatal(args ...interface{}) {
	if logger.CanSentry() {
		logger.SentryException(errors.New(fmt.Sprint(args...)))
	}

	if logger.CanLog() {
		logger.logger.Fatal(args...)
	}
}

func (logger *Logger) Panic(args ...interface{}) {
	if logger.CanSentry() {
		logger.SentryException(errors.New(fmt.Sprint(args...)))
	}

	if logger.CanLog() {
		logger.logger.Panic(args...)
	}
}

func (logger *Logger) Print(v ...interface{}) {
	if logger.CanLog() {
		logger.logger.Print(v...)
	}
}

func (logger *Logger) SentryException(exception error) {
	if logger.CanSentry() {
		logger.sentry.CaptureException(exception)
	}
}

func (logger *Logger) SentryWarn(args ...interface{}) {
	if logger.CanSentry() {
		logger.sentry.CaptureMessage(fmt.Sprint(args...))
	}
}

func (logger *Logger) Logrus() *logrus.Logger {
	return logger.logger
}

func (logger *Logger) CanLog() bool {
	return logger.logger != nil
}

func (logger *Logger) CanSentry() bool {
	return logger.sentry != nil
}

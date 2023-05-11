package rzap

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func NewLogger(config *Config) *Logger {
	logger := config.Build()

	return &Logger{
		logger: logger,
		sugar:  logger.Sugar(),
	}
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.logger.Panic(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}

func (l *Logger) SDebug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *Logger) SDebugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

func (l *Logger) SDebugw(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, keysAndValues...)
}

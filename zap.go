package rzap

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func NewLogger(config *Config) *Logger {
	logger := config.build()

	return &Logger{
		logger: logger,
		sugar:  logger.Sugar(),
	}
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
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

// -----------------

func (l *Logger) SDebug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *Logger) SDebugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

func (l *Logger) SDebugw(msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, keysAndValues...)
}

func (l *Logger) SInfo(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *Logger) SInfof(template string, args ...interface{}) {
	l.sugar.Infof(template, args...)
}

func (l *Logger) SInfow(msg string, keysAndValues ...interface{}) {
	l.sugar.Infow(msg, keysAndValues...)
}

func (l *Logger) SWarn(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *Logger) SWarnf(template string, args ...interface{}) {
	l.sugar.Warnf(template, args...)
}

func (l *Logger) SWarnw(msg string, keysAndValues ...interface{}) {
	l.sugar.Warnw(msg, keysAndValues...)
}

func (l *Logger) SError(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *Logger) SErrorf(template string, args ...interface{}) {
	l.sugar.Errorf(template, args...)
}

func (l *Logger) SErrorw(msg string, keysAndValues ...interface{}) {
	l.sugar.Errorw(msg, keysAndValues...)
}

func (l *Logger) SFatal(args ...interface{}) {
	l.sugar.Fatal(args...)
}

func (l *Logger) SFatalf(template string, args ...interface{}) {
	l.sugar.Fatalf(template, args...)
}

func (l *Logger) SFatalw(msg string, keysAndValues ...interface{}) {
	l.sugar.Fatalw(msg, keysAndValues...)
}

func (l *Logger) SPanic(args ...interface{}) {
	l.sugar.Panic(args...)
}

func (l *Logger) SPanicf(template string, args ...interface{}) {
	l.sugar.Panicf(template, args...)
}

func (l *Logger) SPanicw(msg string, keysAndValues ...interface{}) {
	l.sugar.Panicw(msg, keysAndValues...)
}

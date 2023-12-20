package rzap

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger  *zap.Logger
	sugar   *zap.SugaredLogger
	enabled bool

	Debug func(msg string, fields ...zap.Field)
	Info  func(msg string, fields ...zap.Field)
	Warn  func(msg string, fields ...zap.Field)
	Error func(msg string, fields ...zap.Field)
	Fatal func(msg string, fields ...zap.Field)
	Panic func(msg string, fields ...zap.Field)
}

func NewLogger(config *Config) *Logger {
	logger := config.build()

	lw := &Logger{
		logger:  logger,
		sugar:   logger.Sugar(),
		enabled: config.enabled,
	}

	lw.Debug = lw.makeLogFunc(lw.logger.Debug)
	lw.Info = lw.makeLogFunc(lw.logger.Info)
	lw.Warn = lw.makeLogFunc(lw.logger.Warn)
	lw.Error = lw.makeLogFunc(lw.logger.Error)
	lw.Fatal = lw.makeLogFunc(lw.logger.Fatal)
	lw.Panic = lw.makeLogFunc(lw.logger.Panic)

	return lw
}

func (l *Logger) makeLogFunc(logFunc func(string, ...zap.Field)) func(string, ...zap.Field) {
	return func(msg string, fields ...zap.Field) {
		if l.enabled {
			logFunc(msg, fields...)
		}
	}
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}

/*
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
*/

// -----------------

func (l *Logger) SDebug(args ...interface{}) {
	if l.enabled {
		l.sugar.Debug(args...)
	}
}

func (l *Logger) SDebugf(template string, args ...interface{}) {
	if l.enabled {
		l.sugar.Debugf(template, args...)
	}
}

func (l *Logger) SDebugw(msg string, keysAndValues ...interface{}) {
	if l.enabled {
		l.sugar.Debugw(msg, keysAndValues...)
	}
}

func (l *Logger) SInfo(args ...interface{}) {
	if l.enabled {
		l.sugar.Info(args...)
	}
}

func (l *Logger) SInfof(template string, args ...interface{}) {
	if l.enabled {
		l.sugar.Infof(template, args...)
	}
}

func (l *Logger) SInfow(msg string, keysAndValues ...interface{}) {
	if l.enabled {
		l.sugar.Infow(msg, keysAndValues...)
	}
}

func (l *Logger) SWarn(args ...interface{}) {
	if l.enabled {
		l.sugar.Warn(args...)
	}
}

func (l *Logger) SWarnf(template string, args ...interface{}) {
	if l.enabled {
		l.sugar.Warnf(template, args...)
	}
}

func (l *Logger) SWarnw(msg string, keysAndValues ...interface{}) {
	if l.enabled {
		l.sugar.Warnw(msg, keysAndValues...)
	}
}

func (l *Logger) SError(args ...interface{}) {
	if l.enabled {
		l.sugar.Error(args...)
	}
}

func (l *Logger) SErrorf(template string, args ...interface{}) {
	if l.enabled {
		l.sugar.Errorf(template, args...)
	}
}

func (l *Logger) SErrorw(msg string, keysAndValues ...interface{}) {
	if l.enabled {
		l.sugar.Errorw(msg, keysAndValues...)
	}
}

func (l *Logger) SFatal(args ...interface{}) {
	if l.enabled {
		l.sugar.Fatal(args...)
	}
}

func (l *Logger) SFatalf(template string, args ...interface{}) {
	if l.enabled {
		l.sugar.Fatalf(template, args...)
	}
}

func (l *Logger) SFatalw(msg string, keysAndValues ...interface{}) {
	if l.enabled {
		l.sugar.Fatalw(msg, keysAndValues...)
	}
}

func (l *Logger) SPanic(args ...interface{}) {
	if l.enabled {
		l.sugar.Panic(args...)
	}
}

func (l *Logger) SPanicf(template string, args ...interface{}) {
	if l.enabled {
		l.sugar.Panicf(template, args...)
	}
}

func (l *Logger) SPanicw(msg string, keysAndValues ...interface{}) {
	if l.enabled {
		l.sugar.Panicw(msg, keysAndValues...)
	}
}

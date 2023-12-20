package rzap

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger  *zap.Logger
	sugar   *zap.SugaredLogger
	enabled bool
}

func NewLogger(config *Config) *Logger {
	logger := config.build()

	return &Logger{
		logger:  logger,
		sugar:   logger.Sugar(),
		enabled: config.enabled,
	}
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	if l.enabled {
		l.logger.Debug(msg, fields...)
	}
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	if l.enabled {
		l.logger.Info(msg, fields...)
	}
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	if l.enabled {
		l.logger.Warn(msg, fields...)
	}
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	if l.enabled {
		l.logger.Error(msg, fields...)
	}
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	if l.enabled {
		l.logger.Fatal(msg, fields...)
	}
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	if l.enabled {
		l.logger.Panic(msg, fields...)
	}
}

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

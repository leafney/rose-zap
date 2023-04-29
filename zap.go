package rzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	logger *zap.Logger
}

func New() *Logger {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewJSONEncoder(config)

	writer := zapcore.AddSync(os.Stdout)

	core := zapcore.NewCore(encoder, writer, zap.InfoLevel)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return &Logger{
		logger: logger,
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

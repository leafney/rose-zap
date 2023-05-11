package rzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Config struct {
	//level       string
	level       zapcore.Level
	atomicLevel zap.AtomicLevel
	callerSkip  int
	formatJson  bool
	outConsole  bool
	outFile     *OutFile
}

type OutFile struct {
	enable bool
}

func NewConfig() *Config {
	return &Config{
		//level:       "info",
		atomicLevel: zap.NewAtomicLevel(),
		callerSkip:  1,
	}

}

func (c *Config) build() *zap.Logger {

	encConfig := zapcore.EncoderConfig{
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

	encoder := zapcore.NewJSONEncoder(encConfig)

	writer := zapcore.AddSync(os.Stdout)

	//fmt.Println("level ", conf.level)

	//conf.atomicLevel.SetLevel(getLevel(conf.level))

	core := zapcore.NewCore(encoder, writer, c.atomicLevel)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(c.callerSkip))
}

func (c *Config) SetLevel(level string) *Config {
	c.atomicLevel.SetLevel(getLevel(level))
	return c
}

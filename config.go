package rzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	//level       string
	//level       zapcore.Level
	atomicLevel zap.AtomicLevel
	callerSkip  int
	showCaller  bool
	formatJson  bool
	outConsole  bool
	outType     OTP
	outFile     *FileConfig
}

func NewConfig() *Config {
	return &Config{
		//level:       "info",
		atomicLevel: zap.NewAtomicLevel(),
		callerSkip:  1,
		showCaller:  true,
		outConsole:  true,
		outType:     0,
		formatJson:  true,
		outFile: &FileConfig{
			FileName:   DefaultSingleFilename,
			MaxSize:    1024,
			MaxBackups: 0,
			MaxAge:     1,
			LocalTime:  true,
			Compress:   false,
		},
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

	//encoder := zapcore.NewJSONEncoder(encConfig)
	encoder := GetEncode(c.formatJson, encConfig)

	//writer := zapcore.AddSync(os.Stdout)

	writer := SingleFileWriter(c.outFile, c.outConsole)

	//fmt.Println("level ", conf.level)

	//conf.atomicLevel.SetLevel(getLevel(conf.level))

	core := zapcore.NewCore(encoder, writer, c.atomicLevel)

	return zap.New(core,
		//zap.AddCaller(),
		zap.WithCaller(c.showCaller),
		zap.AddCallerSkip(c.callerSkip),
		zap.AddStacktrace(zapcore.WarnLevel),
	) //
}

func (c *Config) SetLevel(level string) *Config {
	c.atomicLevel.SetLevel(getLevel(level))
	return c
}

func (c *Config) ShowCaller(enabled bool) *Config {
	c.showCaller = enabled
	return c
}

func (c *Config) SetCallSkip(skip int) *Config {
	c.callerSkip = skip
	return c
}

func (c *Config) UseFmtJson(enabled bool) *Config {
	c.formatJson = enabled
	return c
}

//func (c *Config) OutConsole() *Config {
//
//	return c
//}

func (c *Config) OutSingleFileDefault(filename string, outConsole bool) *Config {
	c.outType = OutTypeSingleFileDefault
	c.outConsole = outConsole
	c.outFile.FileName = filename

	return c
}

func (c *Config) OutSingleFileCustom(outConsole bool) *Config {
	c.outType = OutTypeSingleFileCustom
	c.outConsole = outConsole
	return c
}

func (c *Config) OutMultiFilesDefault(outConsole bool) *Config {
	c.outType = OutTypeMultiFileDefault
	c.outConsole = outConsole
	//c.outFile.FileName = filename

	return c
}

func (c *Config) OutMultiFilesCustom(outConsole bool) *Config {
	c.outType = OutTypeMultiFileCustom
	c.outConsole = outConsole
	return c
}

func (c *Config) OutInfoConsoleErrorFile() *Config {
	c.outType = OutTypeInfoError
	c.outConsole = true

	return c
}

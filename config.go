package rzap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	atomicLevel  zap.AtomicLevel
	callerSkip   int
	showCaller   bool
	formatJson   bool
	withConsole  bool
	outType      OTP
	outFile      *FileConfig
	outInfoFile  *FileConfig
	outErrorFile *FileConfig
	outCustom    bool
	//singleFilename string //
	//infoFilename   string
	//errorFilename  string
	showStacktrace bool
}

func NewConfig() *Config {
	return &Config{
		atomicLevel:    zap.NewAtomicLevel(),
		callerSkip:     1,
		showCaller:     true,
		withConsole:    true,
		outType:        OutTypeConsole,
		outCustom:      false,
		formatJson:     true,
		showStacktrace: true,
		outFile: &FileConfig{
			FileName:   DefaultSingleFilename,
			MaxSize:    1024,
			MaxBackups: 0,
			MaxAge:     1,
			LocalTime:  true,
			Compress:   false,
		},
		outInfoFile:  nil,
		outErrorFile: nil,
	}
}

func (c *Config) build() *zap.Logger {

	//encConfig := zap.NewProductionEncoderConfig()

	encConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
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
	//writer := FileWriter(c.outFile, c.withConsole)
	//core := zapcore.NewCore(encoder, writer, c.atomicLevel)

	//fmt.Println("level ", conf.level)
	//conf.atomicLevel.SetLevel(getLevel(conf.level))

	// 多种场景

	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zap.ErrorLevel && level >= zap.DebugLevel
	})

	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})

	consoleWriter := StdOutWriter()
	//infoFileWriter := FileWriter(&FileConfig{}, c.withConsole)
	//errorFileWriter := FileWriter(&FileConfig{}, c.withConsole)

	if c.outCustom {
		if c.outType == OutTypeSingleFileDefault {
			c.outType = OutTypeSingleFileCustom
		} else if c.outType == OutTypeMultiFileDefault {
			c.outType = OutTypeMultiFileCustom
		}
	}

	fmt.Printf("outType [%v]\n", c.outType)
	fmt.Printf("outFile [%v]\n", c.outFile)
	fmt.Printf("outInfoFile [%v]\n", c.outInfoFile)
	fmt.Printf("outErrorFile [%v]\n", c.outErrorFile)

	var core zapcore.Core
	switch c.outType {
	case OutTypeSingleFileDefault:
		fallthrough
	case OutTypeSingleFileCustom:
		singleFileWriter := FileWriter(c.outFile, c.withConsole)
		core = zapcore.NewCore(encoder, singleFileWriter, c.atomicLevel)
	case OutTypeMultiFileDefault:
		fallthrough
	case OutTypeMultiFileCustom:
		infoCfg := c.outFile
		infoCfg.FileName = DefaultMultiFilenameInfo
		if c.outInfoFile != nil {
			infoCfg = c.outInfoFile
		}

		fmt.Printf("infoCfg [%v] outInfoFile [%v]\n", infoCfg, c.outInfoFile)

		infoFileWriter := FileWriter(infoCfg, c.withConsole)
		infoFileCore := zapcore.NewCore(encoder, infoFileWriter, lowPriority)

		errorCfg := c.outFile
		errorCfg.FileName = DefaultMultiFilenameError
		if c.outErrorFile != nil {
			errorCfg = c.outErrorFile
		}

		fmt.Printf("errorCfg [%v] outErrorFile [%v]\n", errorCfg, c.outErrorFile)

		errorFileWriter := FileWriter(errorCfg, c.withConsole)
		errorFileCore := zapcore.NewCore(encoder, errorFileWriter, highPriority)

		core = zapcore.NewTee(infoFileCore, errorFileCore)
	case OutTypeInfoError:
		consoleCore := zapcore.NewCore(encoder, consoleWriter, lowPriority)

		errorCfg := c.outFile
		// Note: It is intentionally designed here that the output error type log file name must be modified through the `SetErrorFileConfig` method
		errorCfg.FileName = DefaultMultiFilenameError
		if c.outErrorFile != nil {
			errorCfg = c.outErrorFile
		}
		errorFileWriter := FileWriter(errorCfg, c.withConsole)
		errorFileCore := zapcore.NewCore(encoder, errorFileWriter, highPriority)

		core = zapcore.NewTee(consoleCore, errorFileCore)
	default:
		core = zapcore.NewCore(encoder, consoleWriter, c.atomicLevel)
	}

	option := make([]zap.Option, 0)
	option = append(option, zap.WithCaller(c.showCaller))
	option = append(option, zap.AddCallerSkip(c.callerSkip))

	if c.showStacktrace {
		option = append(option, zap.AddStacktrace(zapcore.WarnLevel))
	}

	return zap.New(core,
		option...,
	//zap.AddCaller(),
	//zap.WithCaller(c.showCaller),
	//zap.AddCallerSkip(c.callerSkip),
	//zap.AddStacktrace(zapcore.WarnLevel),
	)
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

func (c *Config) ShowStacktrace(enabled bool) *Config {
	c.showStacktrace = enabled
	return c
}

// --------------

/*

// OutSingleFileDefault `filename` default path "logs/rzap.log"
func (c *Config) OutSingleFileDefault(withConsole bool, filename string) *Config {
	c.outType = OutTypeSingleFileDefault
	c.withConsole = withConsole
	if len(filename) == 0 {
		filename = DefaultSingleFilename
	}
	c.singleFilename = filename
	return c
}

func (c *Config) OutSingleFileCustom(withConsole bool) *Config {
	c.outType = OutTypeSingleFileCustom
	c.withConsole = withConsole
	return c
}

// OutMultiFilesDefault `infoFilename` default path "logs/info.log"; `errorFilename` default path "logs/error.log"
func (c *Config) OutMultiFilesDefault(withConsole bool, infoFilename, errorFilename string) *Config {
	c.outType = OutTypeMultiFileDefault
	c.withConsole = withConsole
	if len(infoFilename) == 0 {
		infoFilename = DefaultMultiFilenameInfo
	}
	c.infoFilename = infoFilename
	if len(errorFilename) == 0 {
		errorFilename = DefaultMultiFilenameError
	}
	c.errorFilename = errorFilename
	return c
}

// OutMultiFilesCustom
func (c *Config) OutMultiFilesCustom(withConsole bool) *Config {
	c.outType = OutTypeMultiFileCustom
	c.withConsole = withConsole
	return c
}

// OutInfoConsoleErrorFile `errorFilename` default path "logs/error.log"
func (c *Config) OutInfoConsoleErrorFile(errorFilename string) *Config {
	c.outType = OutTypeInfoError
	c.withConsole = true
	if len(errorFilename) == 0 {
		errorFilename = DefaultMultiFilenameError
	}
	c.errorFilename = errorFilename
	return c
}

*/

// --------------

func (c *Config) OutSingleFile(withConsole bool) *Config {
	c.outType = OutTypeSingleFileDefault
	c.withConsole = withConsole
	return c
}

func (c *Config) OutMultiFile(withConsole bool) *Config {
	c.outType = OutTypeMultiFileDefault
	c.withConsole = withConsole
	return c
}

func (c *Config) OutInfoConsoleErrorFile() *Config {
	c.outType = OutTypeInfoError
	c.withConsole = true
	return c
}

func (c *Config) SetFileConfig(opts ...Option) *Config {
	c.outCustom = true

	for _, opt := range opts {
		opt(c.outFile)
	}
	return c
}

func (c *Config) SetInfoFileConfig(opts ...Option) *Config {
	c.outCustom = true

	c.outInfoFile = new(FileConfig)
	*c.outInfoFile = *c.outFile
	c.outInfoFile.FileName = DefaultMultiFilenameInfo
	for _, opt := range opts {
		opt(c.outInfoFile)
	}
	return c
}

func (c *Config) SetErrorFileConfig(opts ...Option) *Config {
	c.outCustom = true

	c.outErrorFile = new(FileConfig)
	*c.outErrorFile = *c.outFile
	c.outErrorFile.FileName = DefaultMultiFilenameError
	for _, opt := range opts {
		opt(c.outErrorFile)
	}
	return c
}

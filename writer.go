/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     rose-zap
 * @Date:        2023-07-25 17:55
 * @Description:
 */

package rzap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"os"
)

type FileLogConfig struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

func StdOutWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func SingleFileWriter(cfg *FileLogConfig, showStdout bool) zapcore.WriteSyncer {

	fileWriterSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		LocalTime:  true,
		Compress:   false,
	})

	if showStdout {
		return zapcore.NewMultiWriteSyncer(fileWriterSyncer, StdOutWriter())
	}

	return fileWriterSyncer
}

func MultiFilesWriter(showStdout bool) zapcore.WriteSyncer {

	return zapcore.NewMultiWriteSyncer()
}

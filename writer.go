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

func StdOutWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func SingleFileWriter(cfg *FileConfig, showStdout bool) zapcore.WriteSyncer {

	fileWriterSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	})

	if showStdout {
		return zapcore.NewMultiWriteSyncer(fileWriterSyncer, StdOutWriter())
	}

	return fileWriterSyncer
}

func MultiFilesWriter(showStdout bool) zapcore.WriteSyncer {

	return zapcore.NewMultiWriteSyncer()
}

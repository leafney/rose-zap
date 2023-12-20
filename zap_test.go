package rzap

import (
	"errors"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {

	cfg := NewConfig()

	//cfg.SetCallSkip(1)

	cfg.ShowCaller(true).
		UseFmtJson(true).
		//ShowStacktrace(false).
		//OutSingleFileDefault(true, "logs/mylog.log").
		//OutMultiFilesDefault(false, "", "").
		//OutInfoConsoleErrorFile("").
		//OutSingleFile(true).
		OutMultiFile(true).
		SetFileConfig(WithLocalTime(false), WithMaxSize(402), WithFileName("logs/xxx.log")).
		//SetFileConfig(WithFileName(""), WithMaxSize(1), WithMaxBackups(2)).
		//OutInfoConsoleErrorFile().
		//SetFileConfig(WithFileName("logs/cdf.log")).
		SetInfoFileConfig(WithFileName(""), WithMaxSize(33)).
		//SetErrorFileConfig(WithFileName("logs/yyyy.log")).
		SetLevel("debug").SetEnable(true)

	log := NewLogger(cfg)
	defer log.Sync()

	log.Info("测试日志输出", zap.String("name", "tom"))
	log.Debug("调试日志", zap.Int("age", 20))
	log.Info("可以可以")
	log.SErrorf("一个错误 %v", errors.New("不能为空"))
	log.SDebugf("采用 %s 方式", "sugar")
	log.SInfow("日志输出，键值对形式", "attempt", 3, "backoff", time.Second)

}

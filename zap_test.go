package rzap

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {

	config := NewConfig()
	config.SetLevel("debug")
	config.ShowCaller(false)

	log := NewLogger(config)
	defer log.Sync()

	log.Info("测试日志输出", zap.String("name", "tom"))
	log.Debug("调试日志", zap.Int("age", 20))
	log.Info("可以可以")
	log.SDebugf("采用 %s 方式", "sugar")
	log.SInfow("日志输出，键值对形式", "attempt", 3, "backoff", time.Second)
}

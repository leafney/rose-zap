package rzap

import (
	"go.uber.org/zap"
	"testing"
)

func TestNewLogger(t *testing.T) {

	log := New()
	defer log.Sync()

	log.config.SetLevel("debug")

	log.Info("测试日志输出", zap.String("name", "tom"))
	log.Debug("调试日志", zap.Int("age", 20))
	log.Info("可以可以")
}

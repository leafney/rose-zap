package rzap

import "go.uber.org/zap"

type Config struct {
	level       string
	atomicLevel zap.AtomicLevel
}

func newConfig() *Config {
	return &Config{
		level:       "info",
		atomicLevel: zap.NewAtomicLevel(),
	}
}

func (c *Config) SetLevel(level string) {
	c.atomicLevel.SetLevel(getLevel(level))
}

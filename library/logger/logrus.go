package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Logger struct {
	Formatter logrus.Formatter
	Output io.Writer
}

// 默认配置
func DefaultConfig() *Logger {
	return &Logger{
		Formatter: TextFormatter,
		Output: os.Stderr,
	}
}

// build
func (logger *Logger) Build() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(logger.Formatter)
	log.SetOutput(logger.Output)
	return log
}
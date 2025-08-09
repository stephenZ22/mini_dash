package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var mini_log *logrus.Logger

func Init(level string) *logrus.Logger {
	mini_log = logrus.New()
	mini_log.SetOutput(os.Stdout)
	mini_log.SetFormatter(&logrus.JSONFormatter{})

	return mini_log
}

func MiniLogger() *logrus.Logger {
	if mini_log == nil {
		mini_log = Init("info")
	}
	return mini_log
}

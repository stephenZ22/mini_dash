package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var mini_log *logrus.Logger

func Init(level string) *logrus.Logger {
	mini_log = logrus.New()
	mini_log.SetOutput(os.Stdout)
	mini_log.SetFormatter(&logrus.JSONFormatter{})
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Println("Invalid log level, defaulting to info:", err)
		mini_log.SetLevel(logrus.InfoLevel)
		return mini_log
	}

	mini_log.SetLevel(lvl)

	return mini_log
}

func MiniLogger() *logrus.Logger {
	if mini_log == nil {
		mini_log = Init("info")
	}
	return mini_log
}

func GinLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		latency := time.Since(start)

		// 状态码等信息
		status := c.Writer.Status()

		logger.WithFields(logrus.Fields{
			"status":     status,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
		}).Info("HTTP request")
	}
}

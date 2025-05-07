package middleware

import (
	"fmt"
	"os"
	"time"

	"39.com/config"
	"39.com/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 中间件逻辑

		// 调用下一个处理函数
		ctx.Next()
	}
}

// AccessLog 记录访问日志
func AccessLog() gin.HandlerFunc {
	logConf := config.Conf.Log

	if _, err := os.Stat(logConf.Dir); os.IsNotExist(err) {
		os.Mkdir(logConf.Dir, os.ModePerm)
	}

	filename := fmt.Sprintf("%s/%s-%s.log", logConf.Dir, "access", time.Now().Format("2006-01-02"))
	// 配置日志输出到文件
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// 输出到文件
	logrus.SetOutput(file)
	// 设置日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		// 格式化时间
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return func(c *gin.Context) {
		stime := time.Now()
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = utils.GenerateUniqueID("REQ")
		}
		c.Next()
		etime := time.Now()
		// 记录日志
		logrus.WithFields(logrus.Fields{
			"request_id":   requestID,
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"request_url":  c.Request.RequestURI,
			"host":         c.Request.Host,
			"ip":           c.ClientIP(),
			"remote_ip":    c.Request.RemoteAddr,
			"user_agent":   c.Request.UserAgent(),
			"status":       c.Writer.Status(),
			"latency":      etime.Sub(stime).Milliseconds(),
			"time":         stime.Format("2006-01-02 15:04:05"),
			"params":       c.Request.URL.Query(),
			"content_type": c.ContentType(),
		}).Info("web server access log")
	}
}

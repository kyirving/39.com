package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"39.com/config"
	"39.com/internal/dto"
	"39.com/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Apis []string = []string{
	"/user/add",
}

// RequestCacheMiddleware 缓存请求体
func RequestCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 缓存原始请求体
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		bodyBytes, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&buf)

		// 存储到context中供后续使用
		c.Set("raw_body", bodyBytes)
		c.Next()
	}
}

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := &utils.Response{}
		// 获取公共参数
		commonParams, exists := ctx.Get("raw_body")
		if !exists {
			fmt.Println("commonParams:", commonParams)
			resp.Code = utils.PARAMS_ERR
			resp.Msg = "获取公共参数失败"
			resp.Output(ctx)
			ctx.Abort()
			return
		}

		// 公共参数校验
		commonParamsBytes, _ := commonParams.([]byte)
		common := dto.CommonParams{}
		err := json.Unmarshal(commonParamsBytes, &common)
		if err != nil {
			fmt.Println("err:", err)
			resp.Code = utils.PARAMS_ERR
			resp.Output(ctx)
			ctx.Abort()
			return
		}

		// 签名校验
		allParams := make(map[string]interface{})
		if err := json.Unmarshal(commonParamsBytes, &allParams); err != nil {
			fmt.Println("err:", err)
			resp.Code = utils.PARAMS_ERR
			resp.Output(ctx)
			ctx.Abort()
			return
		}

		if allParams["sign"] != utils.Createsign(allParams, config.Conf.App.SecretKey) {
			fmt.Println("sign err:", err)
			resp.Code = utils.UNAUTHORIZED
			resp.Msg = "签名校验失败"
			resp.Output(ctx)
			ctx.Abort()
			return
		}

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
		params, _ := c.Get("raw_body")
		var data map[string]interface{}
		_ = json.Unmarshal(params.([]byte), &data)
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
			"params":       data,
			"content_type": c.ContentType(),
		}).Info("web server access log")
	}
}

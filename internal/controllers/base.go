package controllers

import (
	"fmt"
	"os"
	"time"

	"39.com/config"
	"github.com/sirupsen/logrus"
)

type BaseController struct {
	// 公共属性
	logger *logrus.Logger
}

func NewBaseController() *BaseController {
	logConf := config.Conf.Log

	if _, err := os.Stat(logConf.Dir); os.IsNotExist(err) {
		os.Mkdir(logConf.Dir, os.ModePerm)
	}

	filename := fmt.Sprintf("%s/%s-%s.log", logConf.Dir, "app", time.Now().Format("2006-01-02"))
	// 配置日志输出到文件
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// 创建一个新的日志记录器
	log := logrus.New()
	// 设置日志输出到文件
	log.SetOutput(file)
	log.SetFormatter(&logrus.JSONFormatter{
		// 格式化时间
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &BaseController{
		logger: log,
	}
}

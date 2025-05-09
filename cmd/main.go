package main

import (
	"fmt"

	"39.com/config"
	"39.com/pkg/database"
	"39.com/routes"
	"github.com/jessevdk/go-flags"
)

// 定义一个结构体，用于存储命令行参数
type Options struct {
	Port string `short:"p" long:"port" description:"端口号" required:"true" default:"8080"`
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic info:", err)
		}
	}()

	var options Options
	_, err := flags.Parse(&options)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	// 初始化配置
	config.InitConfig()
	database.InitMysql()

	// 初始化路由
	router := routes.InitRoutes()
	// 启动服务
	router.Run(":" + options.Port)
}

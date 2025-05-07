package main

import (
	"fmt"

	"39.com/config"
	"39.com/pkg/database"
	"39.com/routes"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic info:", err)
		}
	}()
	// 初始化配置
	config.InitConfig()
	database.InitMysql()

	// 初始化路由
	router := routes.InitRoutes()
	// 启动服务
	router.Run(":8080")
}

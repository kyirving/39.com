package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	dir, err := os.Getwd() // 获取当前工作目录
	if err != nil {
		fmt.Println("获取目录失败:", err)
		return
	}

	filePath := dir + "/config"
	// 初始化配置文件
	configName := "production"
	if os.Getenv("APP_ENV") == "dev" {
		configName = "development"
	}
	viper.SetConfigName(configName) // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(filePath)   // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		// 抛出异常
		panic("read config failed:" + err.Error())
	}

	fmt.Println(viper.Get("app.name"))
	fmt.Println("config init success")
}

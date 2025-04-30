package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type MysqlConf struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConf struct {
	Host     string
	Port     int
	Password string
	Database int
}

type AppConf struct {
	Name string
	Env  string
	Port int
}

type Config struct {
	Mysql MysqlConf
	Redis RedisConf
	App   AppConf
}

// 全局配置
var Conf *Config

func GetConfig() *Config {
	return Conf
}

func GetMysqlConf() *MysqlConf {
	return &Conf.Mysql
}

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

	// 解析配置文件
	Conf = &Config{}
	err = viper.Unmarshal(Conf)
	if err != nil {
		panic("unmarshal config failed:" + err.Error())
	}
	// 打印配置文件
	fmt.Println("InitConfig success")
}

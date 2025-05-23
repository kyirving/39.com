package database

import (
	"fmt"
	"time"

	"39.com/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	return db
}

func InitMysql() {
	// 初始化数据库
	conf := config.GetMysqlConf()
	fmt.Println("初始化数据库:", conf)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Database)

	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,   // DSN data source name
		DefaultStringSize:         128,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database :" + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection pool :" + err.Error())
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("数据库连接成功")
}

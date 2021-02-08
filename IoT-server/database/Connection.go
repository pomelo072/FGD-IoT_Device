package database

import (
	"IoT-server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

var Db *gorm.DB

func init() {
	dsn := strings.Join([]string{config.Sysconfig.DBUserName, ":", config.Sysconfig.DBPassword, "@(", config.Sysconfig.DBIp, ":", config.Sysconfig.DBPort, ")/", config.Sysconfig.DBName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	var err error
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 数据库设置
	sqlDB, err := Db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(1 * time.Second)
	// 生成数据表
	Createtable()
}

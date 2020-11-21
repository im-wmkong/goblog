package bootstrap

import (
	"goblog/pkg/model"
	"time"
)

func SetupDB()  {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)
	// 设置每个连接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}
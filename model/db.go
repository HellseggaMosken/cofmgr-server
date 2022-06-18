package model

import (
	"os"
	"time"

	"cofmgr/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
)

// db 数据库连接单例
var db *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase() {
	dsn := os.Getenv("MYSQL_DSN")
	cfg := gorm.Config{}

	if os.Getenv("GORM_LOG_MODE") == "true" {
		cfg.Logger = gl.New(
			logger.Logger(),
			gl.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel: func() gl.LogLevel {
					lv := gl.LogLevel(logger.LoggerLevel())
					if lv > gl.Info {
						return gl.Info
					} else if lv < gl.Silent {
						return gl.Silent
					}
					return lv
				}(), // Log level
				Colorful: true, // 彩色打印
			},
		)
	}

	_db, err := gorm.Open(mysql.Open(dsn), &cfg)
	if err != nil {
		logger.Fatal("数据库连接失败", err)
	}

	// 设置连接池
	if sqlDB, err := _db.DB(); err != nil {
		logger.Fatal("设置数据库连接池失败", err)
	} else {
		// 空闲
		sqlDB.SetMaxIdleConns(50)
		// 打开
		sqlDB.SetMaxOpenConns(100)
		// 超时
		sqlDB.SetConnMaxLifetime(time.Second * 30)
	}

	db = _db
	migrate()
}

func migrate() {
	if err := db.AutoMigrate(
		&User{},
		&Admin{},
		&Conference{},
		&Contribution{},
		&RefereeCof{},
		&RefereeCtb{},
	); err != nil {
		logger.Fatal("数据库迁移失败", err)
	}
}

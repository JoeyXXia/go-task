package database

import (
	"github.com/joey/go-task/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect 连接到数据库并进行初始化
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库模式
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return db, nil
}

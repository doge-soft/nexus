package migrations

import (
	"status-monitor/models"
	"github.com/jinzhu/gorm"
)

// 迁移数据库
func MigrationModels(context *gorm.DB) error {
	context.AutoMigrate(&models.User{})
	return nil
}

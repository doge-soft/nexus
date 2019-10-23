package migrations

import (
	"github.com/jinzhu/gorm"
	"nexus/auth-gateway/models"
)
// 迁移数据库
func MigrationModels(context *gorm.DB) error {
	context.AutoMigrate(&models.User{})
	return nil
}

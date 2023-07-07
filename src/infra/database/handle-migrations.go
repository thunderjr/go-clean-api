package infra_database

import (
	"github.com/thunderjr/go-clean-api/src/infra/database/local/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	models := []any{
		&models.TaskModel{},
		&models.UserModel{},
	}

	return db.AutoMigrate(models...)
}

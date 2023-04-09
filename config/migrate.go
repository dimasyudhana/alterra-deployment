package config

import (
	"github.com/dimasyudhana/latihan-deployment.git/app/features/user/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(repository.User{})
}

package config

import (
	"github.com/dimasyudhana/latihan-deployment.git/app/features/users/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(repository.Users{})
}

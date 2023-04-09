package repository

import (
	"database/sql"

	"github.com/dimasyudhana/latihan-deployment.git/app/features/books/repository"

	"time"
)

type User struct {
	ID               string
	Username         string `gorm:"type:varchar(255)"`
	Phone            string `gorm:"primaryKey;type:varchar(15);unique"`
	Email            string `gorm:"type:varchar(255)"`
	Password         string `gorm:"type:varchar(100)"`
	Confirm_password string `gorm:"type:varchar(100)"`
	Created_at       time.Time
	Updated_at       time.Time
	Deleted_at       sql.NullTime      `gorm:"index"`
	Books            []repository.Book `gorm:"foreignKey:UserID"`
}

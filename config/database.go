package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection(config Configuration) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Name)

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sql, err := db.DB()
	if err != nil {
		return nil, err
	}

	sql.SetMaxIdleConns(5)
	sql.SetMaxOpenConns(50)
	sql.SetConnMaxIdleTime(5 * time.Minute)

	return db, nil
}

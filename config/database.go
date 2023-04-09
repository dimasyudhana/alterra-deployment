package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config AppConfig) (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect database, ", err.Error())
	}

	return db, nil
}

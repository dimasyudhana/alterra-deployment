package repository

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Year      string `json:"year"`
	Publisher string `json:"publisher"`
	UserID    string `json:"id" gorm:"type:varchar(15)"`
}

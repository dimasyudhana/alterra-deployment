package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"time"

	users "github.com/dimasyudhana/latihan-deployment.git/app/features/users"
	"github.com/dimasyudhana/latihan-deployment.git/helper"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.Repository {
	return &UserModel{
		db: db,
	}
}

func (um UserModel) Register(newUser users.Core) (users.Core, error) {
	hashedPassword, err := helper.GenerateHashedPassword(newUser.Password)
	if err != nil {
		log.Error("error while hashing password", err.Error())
		return users.Core{}, err
	}

	inputUser := Users{
		ID:               uuid.New().String(),
		Username:         newUser.Username,
		Phone:            newUser.Phone,
		Email:            newUser.Email,
		Password:         hashedPassword,
		Confirm_password: hashedPassword,
		Created_at:       time.Now(),
		Updated_at:       time.Now(),
		Deleted_at:       sql.NullTime{},
	}

	if err := um.db.Table("users").Create(&inputUser).Error; err != nil {
		log.Error("Error while creating user", err.Error())
		return users.Core{}, err
	}

	return newUser, nil
}

func (um *UserModel) FindByPhone(phone string) ([]users.Core, error) {
	userList := []*Users{}

	if err := um.db.Where("phone = ?", phone).Find(&userList).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return []users.Core{}, nil
		}
		log.Error("Failed to query user by phone", err.Error())
		return nil, err
	}

	var result []users.Core
	for _, u := range userList {
		uc := users.Core{
			Username: u.Username,
			Phone:    u.Phone,
			Password: u.Password,
		}
		result = append(result, uc)
	}

	return result, nil
}

func (um *UserModel) Login(phone string, password string) (users.Core, error) {
	result := Users{}
	//Query Login >> select * from users where hp = ? and password = ?
	if err := um.db.Where("phone = ?", phone).Find(&result).Error; err != nil {
		log.Error(err.Error())
		return users.Core{}, err
	}

	if result.Phone == "" {
		log.Error("Phone tidak ditemukan")
		return users.Core{}, errors.New("phone tidak ditemukan")
	}

	if !helper.CompareHashedPassword(string(result.Password), password) {
		log.Error("Password tidak sesuai")
		return users.Core{}, errors.New("password tidak sesuai")
	}

	return users.Core{Username: result.Username, Phone: result.Phone}, nil
}

func (um *UserModel) UpdateByPhone(phone string, updatedUser users.Core) error {
	user := Users{}
	if err := um.db.Where("phone = ?", phone).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with phone %v not found", phone)
		}
		log.Print("Failed to query user by phone", err)
		return err
	}

	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.Updated_at = time.Now()

	if err := um.db.Save(&user).Error; err != nil {
		log.Print("Failed to update user", err)
		return err
	}

	return nil
}

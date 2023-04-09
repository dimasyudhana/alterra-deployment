package service

import (
	"errors"

	"github.com/dimasyudhana/latihan-deployment.git/app/features/users"
	"github.com/labstack/gommon/log"
)

type UserLogic struct {
	m users.Repository
}

func New(r users.Repository) users.UseCase {
	return &UserLogic{
		m: r,
	}
}

func (ul *UserLogic) Register(newUser users.Core) error {
	_, err := ul.m.Register(newUser)
	if err != nil {
		log.Error("register logic error:", err.Error())
		return errors.New("terjadi kesalahan pada server")
	}
	return nil
}

func (ul *UserLogic) Login(phone string, password string) (users.Core, error) {
	result, err := ul.m.Login(phone, password)
	if err != nil {
		return users.Core{}, errors.New("terdapat permasalahan pada server")
	}
	return result, nil
}

func (ul *UserLogic) UpdateByPhone(phone string, username string, email string) (users.Core, error) {
	_, err := ul.m.FindByPhone(phone)
	if err != nil {
		return users.Core{}, err
	}

	updatedUser := users.Core{
		Phone:    phone,
		Username: username,
		Email:    email,
	}

	if err := ul.m.UpdateByPhone(phone, updatedUser); err != nil {
		return users.Core{}, err
	}

	return updatedUser, nil
}

func (ul *UserLogic) FindByPhone(phone string) ([]users.Core, error) {
	users, err := ul.m.FindByPhone(phone)
	if err != nil {
		log.Error("failed to query user by phone:", err.Error())
		return nil, err
	}
	return users, nil
}

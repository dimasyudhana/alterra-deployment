package usecase

import (
	"errors"

	user "github.com/dimasyudhana/latihan-deployment.git/app/features/user"
)

type UserLogic struct {
	m user.Repository
}

func New(r user.Repository) user.UseCase {
	return &UserLogic{
		m: r,
	}
}

func (ul *UserLogic) Register(newUser user.Core) error {
	_, err := ul.m.Register(newUser)
	if err != nil {
		return errors.New("terjadi kesalahan pada server")
	}

	return nil
}

func (ul *UserLogic) Login(phone string, password string) (user.Core, error) {
	result, err := ul.m.Login(phone, password)
	if err != nil {
		return user.Core{}, errors.New("terdapat permasalahan pada server")
	}
	return result, nil
}

func (ul *UserLogic) UpdateByPhone(phone string, username string, email string) (user.Core, error) {
	_, err := ul.m.FindByPhone(phone)
	if err != nil {
		return user.Core{}, err
	}

	updatedUser := user.Core{
		Phone:    phone,
		Username: username,
		Email:    email,
	}

	if err := ul.m.UpdateByPhone(phone, updatedUser); err != nil {
		return user.Core{}, err
	}

	return updatedUser, nil
}

func (ul *UserLogic) FindByPhone(phone string) ([]*user.Core, error) {
	users, err := ul.m.FindByPhone(phone)
	if err != nil {
		return nil, errors.New("failed to query user by phone")
	}

	return users, nil
}

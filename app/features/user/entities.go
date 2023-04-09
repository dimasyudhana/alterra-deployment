package user

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	Username         string
	Phone            string
	Email            string
	Password         string
	Confirm_password string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type UseCase interface {
	Register(newUser Core) error
	FindByPhone(phone string) ([]*Core, error)
	Login(phone string, password string) (Core, error)
	UpdateByPhone(phone string, username string, email string) (Core, error)
}

type Repository interface {
	Register(newUser Core) (Core, error)
	FindByPhone(phone string) ([]*Core, error)
	Login(phone string, password string) (Core, error)
	UpdateByPhone(phone string, updatedUser Core) error
}

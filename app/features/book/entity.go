package book

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	Title     string
	Year      string
	Publisher string
	Author    string
	User_ID   string
}

type Handler interface {
	InsertBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
}

type UseCase interface {
	InsertBook(newBook Core, user_id string) (Core, error)
	UpdateBook(newBook Core, user_id string) (Core, error)
	GetBookByID(bookID uint, userID string) (Core, error)
}

type Repository interface {
	InsertBook(newBook Core, userID string) (Core, error)
	UpdateBook(updatedBook Core, userID string) (Core, error)
	GetBookByID(bookID uint, userID string) (Core, error)
}

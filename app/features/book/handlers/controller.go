package handlers

import (
	"net/http"
	"strconv"

	"github.com/dimasyudhana/latihan-deployment.git/app/features/book"
	"github.com/dimasyudhana/latihan-deployment.git/helper"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	service book.UseCase
}

func New(s book.UseCase) book.Handler {
	return &BookController{
		service: s,
	}
}

func (bc *BookController) InsertBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := helper.DecodeJWT(c.Get("user").(*jwt.Token))

		input := BookRequest{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "terdapat kesalahan input dari Book", nil))
		}
		res, err := bc.service.InsertBook(book.Core{Title: input.Title, Year: input.Year, Publisher: input.Publisher}, userID)

		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "sukses menambahkan buku", res))
	}
}

func (bc *BookController) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := helper.DecodeJWT(c.Get("user").(*jwt.Token))

		input := BookRequest{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "terdapat kesalahan input dari Book", nil))
		}

		// Get the ID of the book to update from the URL parameter
		bookID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "terdapat kesalahan pada parameter id", nil))
		}

		// Check if the book with the specified ID belongs to the authenticated user
		existingBook, err := bc.service.GetBookByID(uint(bookID), userID)
		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}
		if existingBook.User_ID != userID {
			return c.JSON(helper.ResponseFormat(http.StatusForbidden, "Anda tidak diizinkan untuk memperbarui buku ini", nil))
		}

		// Create a new Core struct with the updated fields
		updatedBook := book.Core{
			Title:     input.Title,
			Year:      input.Year,
			Publisher: input.Publisher,
		}

		// Update the book in the database
		res, err := bc.service.UpdateBook(updatedBook, userID)
		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "sukses memperbarui buku", res))
	}
}

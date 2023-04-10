package usecase

import (
	"errors"

	"github.com/dimasyudhana/latihan-deployment.git/app/features/book"

	"strings"

	"github.com/labstack/gommon/log"
)

type BookModel struct {
	repo book.Repository
}

func New(br book.Repository) book.UseCase {
	return &BookModel{
		repo: br,
	}
}

func (bm *BookModel) InsertBook(newBook book.Core, user_id string) (book.Core, error) {
	result, err := bm.repo.InsertBook(newBook, user_id)
	if err != nil {
		log.Error("terjadi kesalahan input buku", err.Error())
		if strings.Contains(err.Error(), "too much") {
			return book.Core{}, errors.New("terdapat kesalahan input, nilai yang diberikan terlalu panjang")
		}
		return book.Core{}, errors.New("terdapat masalah pada server")
	}
	return result, nil
}

func (bm *BookModel) UpdateBook(newBook book.Core, user_id string) (book.Core, error) {
	// Mengambil buku yang ingin di-update dari database
	oldBook, err := bm.repo.GetBookByID(newBook.ID, user_id)
	if err != nil {
		log.Error("terjadi kesalahan saat mengambil data buku", err.Error())
		return book.Core{}, errors.New("terdapat masalah pada server")
	}

	// Memeriksa apakah buku yang ingin di-update dimiliki oleh user yang sesuai
	if oldBook.User_ID != user_id {
		log.Error("buku yang ingin di-update tidak dimiliki oleh user yang sesuai")
		return book.Core{}, errors.New("anda tidak memiliki akses untuk meng-update buku ini")
	}

	// Mengisi data buku yang ingin di-update
	oldBook.Title = newBook.Title
	oldBook.Publisher = newBook.Publisher
	oldBook.Year = newBook.Year

	// Mengupdate buku di database
	updatedBook, err := bm.repo.UpdateBook(oldBook, user_id)
	if err != nil {
		log.Error("terjadi kesalahan saat meng-update buku", err.Error())
		return book.Core{}, errors.New("terdapat masalah pada server")
	}

	return updatedBook, nil
}

func (bm *BookModel) GetBookByID(bookID uint, userID string) (book.Core, error) {
	book, err := bm.repo.GetBookByID(bookID, userID)
	if err != nil {
		log.Error("terjadi kesalahan saat mengambil data buku", err.Error())
		return book, errors.New("terdapat masalah pada server")
	}

	return book, nil
}

package repository

import (
	"log"

	"github.com/dimasyudhana/latihan-deployment.git/app/features/book"

	"gorm.io/gorm"
)

type BookModel struct {
	db *gorm.DB
}

func New(d *gorm.DB) book.Repository {
	return &BookModel{
		db: d,
	}
}

func (bm *BookModel) InsertBook(newBook book.Core, userID string) (book.Core, error) {
	// Membuat daftar buku baru
	var insertBook Book
	insertBook.Title = newBook.Title
	insertBook.Publisher = newBook.Publisher
	insertBook.Year = newBook.Year
	insertBook.UserID = userID

	if err := bm.db.Table("books").Create(&insertBook).Error; err != nil {
		log.Println("Terjadi error saat membuat daftar buku baru.", err.Error())
		return book.Core{}, err
	}
	return newBook, nil
}

func (bm *BookModel) UpdateBook(updatedBook book.Core, userID string) (book.Core, error) {
	// Mencari buku yang akan diupdate di database
	var existingBook Book
	if err := bm.db.Table("books").Where("id = ? AND user_id = ?", updatedBook.ID, userID).First(&existingBook).Error; err != nil {
		log.Println("Terjadi error saat mencari buku yang akan diupdate.", err.Error())
		return book.Core{}, err
	}

	// Mengisi data baru ke dalam buku yang ditemukan
	existingBook.Title = updatedBook.Title
	existingBook.Publisher = updatedBook.Publisher
	existingBook.Year = updatedBook.Year

	// Melakukan update buku di database
	if err := bm.db.Table("books").Save(&existingBook).Error; err != nil {
		log.Println("Terjadi error saat melakukan update buku.", err.Error())
		return book.Core{}, err
	}

	return existingBook.Book(), nil
}

func (b *Book) Book() book.Core {
	return book.Core{
		ID:        b.ID,
		Title:     b.Title,
		Year:      b.Year,
		Publisher: b.Publisher,
	}
}

func (bm *BookModel) GetBookByID(bookID uint, userID string) (book.Core, error) {
	var bookData Book

	// Mengambil data buku dari database dengan ID dan user ID tertentu
	if err := bm.db.Table("books").Where("id = ? AND user_id = ?", bookID, userID).First(&bookData).Error; err != nil {
		log.Println("Terjadi error saat mengambil data buku dari database.", err.Error())
		return book.Core{}, err
	}

	return bookData.Book(), nil
}

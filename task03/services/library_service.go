package services

import (
	"task03/models"
)

type LibraryManaager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int , memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

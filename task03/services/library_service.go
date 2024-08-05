package services

import (
	"task03/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int , memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	allBooks 	map[int]models.Book
	allMembers 	map[int]models.Member
}

func AddBook(book models.Book) {

}

func RemoveBook(bookID int) {

}
func BorrowBook(bookID int, memberID int) error {
	return nil
}

func ReturnBook(bookID int, memberID int) error {
	return nil
}
func ListAvailableBooks() []models.Book {
	return nil
}
func ListBorrowedBooks() []models.Book {
	return nil
}

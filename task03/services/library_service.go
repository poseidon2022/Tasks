package services

import (
	"errors"
	"fmt"
	"task03/models"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int , memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books 	    map[int]models.Book
	Members 	map[int]models.Member
}

func (l *Library) AddBook(book models.Book) error {
	for index, _ := range l.Books{
		if index == book.ID {
			return errors.New("book ID already exists")
		}
	}
	l.Books[book.ID] = book
	fmt.Println("Book Added Successfully")
	return nil
}

func (l *Library) RemoveBook(bookID int) error {
	for index, _ := range l.Books{
		if index == bookID {
			delete(l.Books, index)
			fmt.Println("Book removed successfully!")
			return nil
		}
	}
	return errors.New("book is not in the library")
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bOk := l.Books[bookID]
	member, mOk := l.Members[memberID]

	if !bOk {
		return errors.New("a book with the given id does not exist")
	}
	if !mOk {
		return errors.New("a member with the given id does not exist")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	fmt.Println("Book borrowed successfully!")
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, bOk := l.Books[bookID]
	member, mOk := l.Members[memberID]

	if !bOk {
		return errors.New("a book with the given id does not exist")
	}
	if !mOk {
		return errors.New("a member with the given id does not exist")
	}
	if book.Status == "Available" {
		return errors.New("book is available")
	}

	book.Status = "Available"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	fmt.Println("Book returned successfully!")
	return nil
}
func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := []models.Book{}
	for _, value := range l.Books {
		if value.Status == "Available" {
			availableBooks = append(availableBooks, value)
		}
	}

	return availableBooks
}
func (l *Library) ListBorrowedBooks() []models.Book {
	borrowedBooks := []models.Book{}
	for _, value := range l.Books {
		if value.Status == "Available" {
			borrowedBooks = append(borrowedBooks, value)
		}
	}

	return borrowedBooks
}

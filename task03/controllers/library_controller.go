package controllers

import (
	"task03/services"
	"task03/models"
	"strconv"
	"os"
	"bufio"
	"fmt"
	"strings"
)

type Library = services.Library

func BookInput(l *Library) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the ID of the book to be added: ")
	book_id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while accepting the ID")
		return
	}

	book_id = strings.TrimSpace(book_id)
	book_id_num ,err := strconv.Atoi(book_id)
	for err != nil {
		fmt.Println("Please enter a valid id number")
		fmt.Print("Enter the ID of the book to be added: ")
		book_id, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while accepting the ID")
			return
		}
		book_id = strings.TrimSpace(book_id)
		book_id_num , err = strconv.Atoi(book_id)
	}

	fmt.Print("Enter the Title of the book to be added: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while accepting the ID")
		return
	}

	title = strings.TrimSpace(title)

	fmt.Print("Enter the author of the book to be added: ")
	author, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while accepting the ID")
		return
	}

	author = strings.TrimSpace(author)
	
	status := "Available"

	newBook := models.Book{
		ID : book_id_num, 
		Title: title,
		Author: author,
		Status: status,
	}

	fmt.Println(l.AddBook(newBook))
}

func BookRemove(l *Library) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the ID of the book you want to remove: ")
	book_id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while accepting the ID")
		return
	}

	book_id = strings.TrimSpace(book_id)
	book_id_num ,err := strconv.Atoi(book_id)
	for err != nil {
		fmt.Println("Please enter a valid id number.")
		fmt.Print("Enter the ID of the book you want to remove: ")
		book_id, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while accepting the ID")
			return
		}
		book_id = strings.TrimSpace(book_id)
		book_id_num, err = strconv.Atoi(book_id)
	}
	fmt.Println(l.RemoveBook(book_id_num))
}

func BookBorrow(l *Library) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the ID of the book you want to borrow: ")
	book_id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while accepting the ID")
		return
	}

	book_id = strings.TrimSpace(book_id)
	book_id_num ,err := strconv.Atoi(book_id)
	for err != nil {
		fmt.Println("Please enter a valid id number")
		fmt.Print("Enter the ID of the book you want to borrow: ")
		book_id, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while accepting the ID")
			return
		}
		book_id = strings.TrimSpace(book_id)
		book_id_num , err = strconv.Atoi(book_id)
	}

	fmt.Print("Enter your membership id number: ")
	member_id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while accepting the ID")
		return
	}

	member_id = strings.TrimSpace(member_id)
	member_id_num ,err := strconv.Atoi(member_id)
	for err != nil {
		fmt.Println("Please enter a valid id number")
		fmt.Print("Enter your membership id number: ")
		member_id, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while accepting the ID")
			return
		}
		member_id = strings.TrimSpace(member_id)
		member_id_num , err = strconv.Atoi(member_id)
	}

	fmt.Println(l.BorrowBook(book_id_num, member_id_num))
}

func BookReturn(l *Library) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the ID of the book you want to return: ")
	book_id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while accepting the ID")
		return
	}

	book_id = strings.TrimSpace(book_id)
	book_id_num ,err := strconv.Atoi(book_id)
	for err != nil {
		fmt.Println("Please enter a valid id number")
		fmt.Print("Enter the ID of the book you want to return: ")
		book_id, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while accepting the ID")
			return
		}
		book_id = strings.TrimSpace(book_id)
		book_id_num , err = strconv.Atoi(book_id)
	}

	fmt.Print("Enter your membership id number: ")
	member_id, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while accepting the ID")
		return
	}

	member_id = strings.TrimSpace(member_id)
	member_id_num ,err := strconv.Atoi(member_id)
	for err != nil {
		fmt.Println("Please enter a valid id number")
		fmt.Print("Enter your membership id number: ")
		member_id, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while accepting the ID")
			return
		}
		member_id = strings.TrimSpace(member_id)
		member_id_num , err = strconv.Atoi(member_id)
	}
	
	fmt.Println(l.ReturnBook(book_id_num, member_id_num))
}

func AvailableBooks(l *Library) {
    fmt.Println(l.ListAvailableBooks())
}
func BorrowedBooks(l *Library) {
	fmt.Println(l.ListBorrowedBooks())
}
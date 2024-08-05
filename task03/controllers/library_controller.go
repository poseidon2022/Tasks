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



func bookInput() {
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
	services.AddBook(newBook)
}

func bookRemove() {

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

	services.RemoveBook(book_id_num)
}

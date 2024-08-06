package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	controller "task03/controllers"
	"task03/models"
	"task03/services"
)

func menu() {
	fmt.Println("Library Management System")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")
	fmt.Println("7. Exit")
}



func main() {
	var Library = services.Library{
		Books : map[int]models.Book{},
		Members: map[int]models.Member{
			1 : {
				ID: 1,
				Name: "Kidus Melaku",
				BorrowedBooks: []models.Book{},
			},
			2 : {
				ID: 2,
				Name: "Bisrat Berhanu",
				BorrowedBooks: []models.Book{},
			},
			3 : {
				ID: 3,
				Name: "Yonatan Tizazu",
				BorrowedBooks: []models.Book{},
			},
			4 : {
				ID: 4,
				Name: "Yohannes Solomon",
				BorrowedBooks: []models.Book{},
			},
		},
	}
	menu()

	fmt.Print("Enter the command you want to execute: ")
	for {
	    reader := bufio.NewReader(os.Stdin)
	    userChoice, err := reader.ReadString('\n')
	    if err != nil {
		    fmt.Println("Error reading user input")
		    return
	    }

		userChoice = strings.TrimSpace(userChoice)
		switch userChoice {
	        case "1":
                controller.BookInput(&Library)
			case "2":
                controller.BookRemove(&Library)
			case "3":
                controller.BookBorrow(&Library)
			case "4":
                controller.BookReturn(&Library)
			case "5":
                controller.AvailableBooks(&Library)
			case "6":
                controller.BorrowedBooks(&Library)
			case "7":
				return
			default:
                fmt.Println("Please enter a valid command")
		}
}
}
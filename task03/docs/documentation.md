## Overview

This Library Management System is a command-line application built using the Go programming language. It allows users to manage books and members within a library. The system provides functionalities for adding, removing, borrowing, and returning books, as well as listing available and borrowed books.

## Project Structure

The project is organized into the following packages:

## controllers

Contains functions for handling user commands.
models: Defines the data structures used in the application.
services: Implements the core logic for managing the library.


## Services

Library

The Library struct contains the core data structures for managing books and members. It includes methods to add, remove, borrow, and return books.

BorrowBook

The BorrowBook method allows a member to borrow a book from the library. It checks if the book exists, if the member exists, and if the book is not already borrowed. If all conditions are met, it updates the book's status to "Borrowed" and adds the book to the member's list of borrowed books.

Controllers

The controllers package contains functions that handle user commands and interact with the Library service.

* BookInput: Handles the command to add a book to the library.

* BookRemove: Handles the command to remove a book from the library.

* BookBorrow: Handles the command to borrow a book from the library. It calls the BorrowBook method of the Library service.

* BookReturn: Handles the command to return a book to the library.

* AvailableBooks: Lists all books that are available for borrowing.

* BorrowedBooks: Lists all books that have been borrowed by members.

## Main Function

The main function initializes the library with some members and sets up a command-line interface to interact with the library. It presents a menu to the user and handles user input to execute the appropriate commands.

### Menu Options

* Add Book: Adds a new book to the library.
* Remove Book: Removes a book from the library.
* Borrow Book: Borrows a book from the library.
* Return Book: Returns a borrowed book to the library.
* List Available Books: Lists all available books in the library.
* List Borrowed Books: Lists all borrowed books in the library.
* Exit: Exits the application.

## Example Usage

* The user is prompted to enter a command to execute.
* Depending on the user's choice, the corresponding function from the controllers package is called.
* For example, if the user chooses to borrow a book, the BookBorrow function is called, which interacts with the Library service to borrow the book.

## Error Handling
The application includes error handling to manage cases where a book or member does not exist, or a book is already borrowed. Appropriate error messages are displayed to the user.
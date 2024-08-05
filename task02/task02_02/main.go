package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func reverseString(input string) string {
	var reversedString string
	i := len(input)-1
	for i > -1 {
		reversedString += string(input[i])
		i --
	}
	return reversedString
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a sentence to check if it is a palindrome or not: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("There was an error while reading your input")
		return
	}
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	var inputModified string
	for _, chr := range input {
		if int(chr) >= 97 && int(chr) <= 123 {
			inputModified += string(chr)
		}
	}

	reversedInput := reverseString(inputModified)
	if reversedInput == inputModified {
		fmt.Println("The word you entered indeed is a plaindrome")
		return
	}

	fmt.Println("The word you entered is not a palindrome.")
}
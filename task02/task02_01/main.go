package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string to count the number of words: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Please enter valid set of characyters.")
		return
	}
	input = strings.TrimSpace(input)
	wordSlice := strings.Split(input,  " ")

	wordCount := map[string]int{} 
	for _, word := range wordSlice {
		var newWord string 
		word = strings.ToLower(word)
		for _, chr := range word {
			if int(chr) >= 97 && int(chr) <= 123 {
				newWord +=string(chr)
			}
		}
		wordCount[newWord] += 1
	}
	fmt.Println("--------Here is the count of each word-------")
	for index,value := range wordCount {
		fmt.Printf("%-10v %v\n", index+":", value)
	}
}
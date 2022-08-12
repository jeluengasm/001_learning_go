package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) bool {
	text = strings.ToLower(text)
	var textReversed string

	for i := len(text) - 1; i >= 0; i-- {
		textReversed += string(text[i])
	}

	if textReversed == text {
		return true
	} else {
		return false
	}
}

func main() {
	slice := []string{"Hello", "World", "from", "Golang"}

	for key, val := range slice {
		fmt.Println(key, val)
	}

	fmt.Println()

	for _, val := range slice {
		fmt.Println(val)
	}

	fmt.Println()

	for key := range slice {
		fmt.Println(key)
	}

	fmt.Println("Ama:", isPalindrome("Ama"))
	fmt.Println("taco cat:", isPalindrome("taco cat"))
}

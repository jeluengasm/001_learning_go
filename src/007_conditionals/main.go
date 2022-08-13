package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	value1 := 1
	value2 := 2

	if value1 == 1 {
		fmt.Println("It's 1")
	} else {
		fmt.Println("Isn't 1")
	}

	if value1 == 1 && value2 == 2 {
		fmt.Println("It's true AND")
	}

	if value1 == 0 || value2 == 2 {
		fmt.Println("It's true OR")
	}

	val, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", val)
}

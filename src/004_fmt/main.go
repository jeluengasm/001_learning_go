package main

import "fmt"

func main() {
	// Classsic Println
	// Variables
	helloMsg := "Hello"
	worldMsg := "World"

	fmt.Println(helloMsg, worldMsg)
	fmt.Println(helloMsg, worldMsg)

	// Printf
	name := "Jhon"
	quantity := 100

	fmt.Printf("%s has finished more than %d courses.\n", name, quantity)
	fmt.Printf("%v has finished more than %v courses.\n", name, quantity)

	// Sprintf
	msg := fmt.Sprintf("%s has finished more than %d courses.", name, quantity)
	fmt.Println(msg)

	// Get datatype
	fmt.Printf("helloMsg: %T", helloMsg)
	fmt.Printf("\nworldMsg: %T", worldMsg)
}

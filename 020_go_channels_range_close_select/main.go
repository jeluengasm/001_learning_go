package main

import (
	"fmt"
)

func message(text string, c chan<- string) { // Only input channel
	c <- text
}

func main() {
	c := make(chan string, 2) // channel with 1 routine (static)
	c <- "Message 1"
	// c <- "Message 2"

	// Len -> How many goroutines, cap -> Maximum capacity of the channel
	fmt.Println("C -> ", "len:", len(c), "cap", cap(c))

	c2 := make(chan string, 2) // channel with 2 routine (static). Can be dynamically allocated
	c2 <- "Message 1"
	c2 <- "Message 2"

	fmt.Println("C2 -> ", "len:", len(c2), "cap", cap(c2))

	//Range -> Iterate over goroutines, close -> Close the channel and don't let add future goroutines
	close(c)
	// c <- "Message 2" // raises error
	close(c2)

	for msg := range c2 {
		fmt.Println(msg) // Requires close the channel before
	}

	// Select
	fmt.Println("\nSelect structure:")
	fmt.Println()
	email1 := make(chan string) // If doesn't specify the number channels, it will be dinammically allocated
	email2 := make(chan string) // If doesn't specify the number channels, it will be dinammically allocated

	go message("Message 1", email1)
	go message("Message 2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email sended: ", m1)
		case m2 := <-email2:
			fmt.Println("Email sended: ", m2)
		}
	}
}

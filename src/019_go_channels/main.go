package main

import "fmt"

func say(s string, c chan<- string) { // Only input channel
	c <- s
}

func main() {
	c := make(chan string, 1) // channel with 1 routine (static)
	fmt.Println("Hello")
	go say("Bye", c)

	fmt.Println(<-c) // Get response from the channel (output channel)
}

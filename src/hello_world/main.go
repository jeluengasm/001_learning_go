package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hello") // execute this command if the currently function was finished
	fmt.Println("World")

	// Continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		// Continue
		if i == 2 {
			fmt.Println("Continue")
			continue
		}

		// break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

}

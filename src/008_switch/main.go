package main

import "fmt"

func main() {
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}

	switch modulo2 := 4 % 2; modulo2 {
	case 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}

	value2 := 200
	switch {
	case value2 > 100:
		fmt.Println("Is greather than 100")
	case value2 < 0:
		fmt.Println("Is less than 0")
	default:
		fmt.Println("Undefined")

	}
}

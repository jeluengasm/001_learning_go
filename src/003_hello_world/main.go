package main

import "fmt"

func main() {

	// Calculating area

	const baseSquare = 10
	squareArea := baseSquare * baseSquare
	fmt.Println("Square area:", squareArea)

	x := 10
	y := 50

	// Sum
	result := x + y
	fmt.Println("Sum:", result)

	// Substract
	result = x - y
	fmt.Println("Substract:", result)

	// Product
	result = x * y
	fmt.Println("Product:", result)

	// Divission
	result = x / y
	fmt.Println("Divission:", result)

	// Modulo
	result = x % y
	fmt.Println("Modulo:", result)

	// Increment
	x++
	fmt.Println("Increment:", x)

	// Decrement
	x--
	fmt.Println("Decrement:", x)
}

package main

import "fmt"

func main() {

	// Constants

	const pi float64 = 3.14

	const pi2 = 3.13

	fmt.Println("pi:", pi, "\npi2:", pi2)

	// Variables

	base := 12
	var height int = 14
	var area int

	fmt.Println(base, height, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calculating area

	const baseSquare = 10
	squareArea := baseSquare * baseSquare
	fmt.Println("Square area:", squareArea)
}

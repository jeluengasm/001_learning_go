package main

import "fmt"

// Interface definition
type figure2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return s.base * s.base
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

// interface function calculation
func calculate(f figure2D) {
	fmt.Println("Area: ", f.area())
}

func main() {

	// interface
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 4}

	calculate(mySquare)
	calculate(myRectangle)

	// List of interfaces

	myInterfaces := []interface{}{"hello", 12, 1.0}
	fmt.Println(myInterfaces...)
}

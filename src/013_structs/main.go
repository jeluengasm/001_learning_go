package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{
		brand: "Ford",
		year:  2020,
	}
	fmt.Println(myCar)

	// Another way
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
	otherCar.year = 2023
	fmt.Println(otherCar)

}

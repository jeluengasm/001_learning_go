package main

import (
	"fmt"
	pk "hello_world/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	pk.PrintMessage("Hello world!")

	// pk.printMessage("Hello world private!")
}

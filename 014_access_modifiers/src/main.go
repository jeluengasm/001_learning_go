package main

import (
	"fmt"

	pk "github.com/jeluengasm/001_learning_go/014_access_modifiers/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	pk.PrintMessage("Hello world!")

	// pk.printMessage("Hello world private!")
}

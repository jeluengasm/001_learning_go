package main

import (
	pc "001_learning_go/src/016_stringers/src/pc"
	"fmt"
)

func main() {
	myPc := pc.New(16, 100, "msi")

	fmt.Println(myPc)
}

package main

import (
	"fmt"

	pc "github.com/jeluengasm/001_learning_go/016_stringers/src/pc"
)

func main() {
	myPc := pc.New(16, 100, "msi")

	fmt.Println(myPc)
}

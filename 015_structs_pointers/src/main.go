package main

import (
	"fmt"

	pc "github.com/jeluengasm/001_learning_go/015_structs_pointers/src/pc"
)

func main() {
	a := 50
	b := &a // Reference

	fmt.Println(b)
	fmt.Println(*b) // Pointer

	*b = 100
	fmt.Println(a)

	// myPC := Pc{
	// 	ram:   16,
	// 	disk:  200,
	// 	brand: "msi",
	// }

	// fmt.Println(myPC)
	// myPC.ping()

	myPc := pc.New(12, 200, "HP")
	myPc.SetRam(16)
	myPc.Describe()
	fmt.Println("\nRAM duplicated")
	myPc.DuplicateRAM()
	myPc.Describe()
}

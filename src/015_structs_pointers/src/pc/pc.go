package pc

import "fmt"

type Pc struct {
	ram   int
	disk  int
	brand string
}

func New(ram, disk int, brand string) Pc {
	myPc := Pc{ram: ram, disk: disk, brand: brand}
	return myPc
}

func (myPc Pc) Describe() {
	fmt.Printf(
		"Brand: %s \nRAM: %dGB\nDisk size: %dGB.\n",
		myPc.brand,
		myPc.ram,
		myPc.disk,
	)
}

func (myPC Pc) Ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *Pc) DuplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func (myPc Pc) GetRam() int {
	return myPc.ram
}

func (myPc *Pc) SetRam(rm int) {
	myPc.ram = rm
}

package main

import "fmt"

func getMessage(s1 string, s2 string) {
	fmt.Println(s1, s2)
}

func duplicateNumber(n int32) int32 {
	return n * 2
}

func doubleReturn(a int) (b, c int) {
	return a, a * 2
}

func main() {
	getMessage("Hello,", "world!")
	val1 := duplicateNumber(4)
	fmt.Println(val1)
	val2, val3 := doubleReturn(4)
	fmt.Println(val2, val3)

	val4, _ := doubleReturn(4) // if one of the values wont need it
	fmt.Println("Val4:", val4)
}

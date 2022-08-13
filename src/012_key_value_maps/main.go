package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["hello"] = 10
	m["world"] = 14

	fmt.Println(m)

	// Slicing maps
	for key, val := range m {
		fmt.Println(key, val)
	}

	value := m["hello"]
	fmt.Println(value)

	value2 := m["hellooo"] // If doesn't exist, return zero value
	fmt.Println(value2)

	value2, confirm := m["hellooo"] // If doesn't exist, return zero value
	fmt.Println(value2, confirm)
}

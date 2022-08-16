package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup // better than time.Sleep

	fmt.Println("Hello")
	wg.Add(1)

	go say("world", &wg)
	wg.Wait()

	go func(s string) {
		fmt.Println(s)
	}("Bye")
	time.Sleep(1 * time.Second) // Bad practice

}

package main

import (
	"fmt"
	"time"
)

func main() {
	// direct call
	fun("direct call")

	// goroutine function call
	go fun("goroutine-1")

	// goroutine with anonymous functions
	go func() {
		fun("goroutine-2")
	}()

	// goroutine with function value call
	fv := fun
	go fv("goroutine-3")

	fmt.Println("waiting for goroutines...")
	time.Sleep(1 * time.Second)
	fmt.Println("done !!!")
}

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

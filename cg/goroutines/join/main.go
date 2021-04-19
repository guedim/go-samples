package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var data int
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		data++
	}()

	wg.Wait()
	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}

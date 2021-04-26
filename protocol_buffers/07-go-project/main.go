package main

import (
	"fmt"

	"github.com/guedim/07-go-project/src/simple"
)

func main() {
	fmt.Println("Hello World")
	doSimple()
}

func doSimple() {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Mario",
		SampleList: []int32{1, 2, 3, 4, 5},
	}

	fmt.Println(sm)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	fmt.Println(arg1, arg2)
}

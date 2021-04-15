package main

import "fmt"

func main() {
	colors01 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors01)

	var colors02 map[string]string
	fmt.Println(colors02)

	colors03 := make(map[string]string)
	colors03["white"] = "#ffffff"
	colors03["black"] = "#000000"
	fmt.Println(colors03)

	delete(colors03, "white")
	fmt.Println(colors03)
}

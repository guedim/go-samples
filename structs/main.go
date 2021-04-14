package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	matias := person{"Matías", "Guerrero"}
	fmt.Println(matias)

	yuly := person{firstName: "Yuly", lastName: "Fuentes"}
	fmt.Println(yuly)

	var mario person

	mario.firstName = "Mario"
	mario.lastName = "Guerrero"
	fmt.Printf("%+v", mario)

}

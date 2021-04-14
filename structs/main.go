package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	matias := person{
		firstName: "Matías",
		lastName:  "Guerrero",
		contact:   contactInfo{"matias@gmail.com", 111111},
	}
	fmt.Println(matias)

	yuly := person{
		firstName: "Yuly",
		lastName:  "Fuentes",
		contact: contactInfo{
			email:   "yulyfuentes1@gmail.com",
			zipCode: 111111,
		}}
	fmt.Println(yuly)

	var mario person

	mario.firstName = "Mario"
	mario.lastName = "Guerrero"
	mario.contact.email = "guedim@gmail.com"
	mario.contact.zipCode = 111111
	fmt.Printf("%+v", mario)

}

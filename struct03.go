package main 

import (
	"fmt"
)

type UserAddress struct {
	street 	string
	city 	string
	state 	string
	zip 	string
}

type User struct {
	name 	string
	age 	int
	address UserAddress
}

func main() {
	user := &User{
		name: "Bob",
		age: 35,
		address: UserAddress{
			street:	"789 Broadway",
			city: 	"New York",
			state: 	"NY",
			zip: 	"10003",
		},
	}

	fmt.Println(user)
	fmt.Println("Name: ", user.name, "\nAge: ", user.age, "\nAddress: ", user.address, "\n")
}
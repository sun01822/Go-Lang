package main

import (
	"fmt"
)

type Player struct {
	Name string // Exported 
	Age int 
	city string // unexported
	state string 
}


type Address struct {
	street string
	city string
	state string
	zip string
}

type Person1 struct {
	name string
	age int
	address Address
}



func main() {
	var player1 Player
	//set the value for player1 struct/object property
	player1.Name = "Sun"
	player1.Age = 24
	player1.city = "Rajshahi"
	player1.state = "Aligonj"

	fmt.Println(player1)

	player2 := Player{"Moon", 21, "Chapai Nawabganj", "Ranihati"}
	fmt.Println(player2)


	person1 := Person1{
        name: "Bob",
        age:  35,
        address: Address{
            street: "789 Broadway",
            city:   "New York",
            state:  "NY",
            zip:    "10003",
        },
    }
	fmt.Println(person1)

	person2 := &Person1{
		name: "John",
		age: 22,
		address: Address{
			street: "759 Alianway",
			city: "London",
			state: "LN",
			zip: "11002",
		},
	}
	fmt.Println(person2)
}
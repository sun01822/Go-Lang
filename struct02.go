package main

import (
	"fmt"
)

type Addresses struct {
	village string
	thana string
	zilla string
}

type Person2 struct {
	name string
	age int
	address []Addresses  // add all addresses using array
}


func main() {
	person1 := Person2{
		name: "Sun",
		age: 24,
		address: []Addresses{
			{
				village: "Ranihati",
				thana: "Shibganj",
				zilla: "Chapai Nawabganj",
			},
			{
				village: "Aligonj",
				thana: "Paba",
				zilla: "Rajshahi",
			},
		},
	}
	fmt.Println(person1)
}
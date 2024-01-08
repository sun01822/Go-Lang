package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Sun", Age: 25}
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Create file success!")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	file.Close()

	file2, err := os.Open("person.json")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		decoder := json.NewDecoder(file2)
		var person2 Person
		err = decoder.Decode(&person2)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Decode success!")
			fmt.Println(person2)
		}
	}
	file.Close()
}

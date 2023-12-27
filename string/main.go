package main

import (
	"fmt"
	"strings"
	
)

func main() {
	// string example
	var myString string = "This is a string"
	fmt.Println(myString)

	// concatenation
	str1 := "Hello, "
	str2 := "World!"
	result := str1 + str2
	fmt.Println(result)

	// slicing
	fmt.Println(result[0:5])

	// indexing
	fmt.Printf("%c\n", result[7])


	// length of a string
	fmt.Println(len(result))

	// convert to upper case
	fmt.Println(strings.ToUpper(result))


	// convert to lower case
	fmt.Println(strings.ToLower(result))
	

	str := "one,two,three,four,five"
	splitString := strings.Split(str, ",")
	fmt.Println(splitString)
	fmt.Println(splitString[2])


	// boolean
	var bVal bool // default is false
	fmt.Printf("bVal: %v\n", bVal)
	
}
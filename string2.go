package main

import (
	"fmt"
	"strings"
)

func main() {
	//	string expample
	var myString string = "Hello World!"
	fmt.Println(myString)

	//	concatenatinon
	str1 := "Hello, "
	str2 := "World!"

	result := str1+str2
	fmt.Println(result)   //	Output: Hello, World!


	//	slicing 
	myString2 := "Hello, World!"
	fmt.Println(myString2[0:5])   //	Output: Hello

	//	indexing
	fmt.Println(myString2[7]) 	//	Output: W

	//	length of a string
	fmt.Println(len(myString2))	//	Output: 13


	//	converting a string to uppercase
	fmt.Println(strings.ToUpper(myString2))	//	Output: HELLO, WORLD!


	// converting a string to lowercase
	fmt.Println(strings.ToLower(myString2))	//	Output: hello, world!


	// spliting a string
	str := "one,two,three,four,five"
	splitStr := strings.Split(str, ",")
	fmt.Println(splitStr)	//	Output: [one two three four five]
	fmt.Println(splitStr[0])	//	Output:	


}
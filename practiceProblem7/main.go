package main

import (
	"fmt"
)

func PrintValues(slice []interface{}) {
	for _, value := range slice {
		fmt.Println(value)
	}
}

func main(){
	values := []interface{}{1, "Hello", 3.14, true, 'A'}

	fmt.Println("PrintValues: ")
	PrintValues(values)
}

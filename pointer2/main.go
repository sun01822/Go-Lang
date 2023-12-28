package main

import "fmt"

func main() {
	a := 4
	b := &a
	// Using address of operator(&) and
    // pointer indirection(*) operator
	fmt.Println(a, b)
	fmt.Println(*b)
	*b = 5
	fmt.Println(a)
}
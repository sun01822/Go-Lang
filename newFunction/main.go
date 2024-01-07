package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	p := new(int) // allocate memory for an int using new() function

	// Before assigning any value
	fmt.Println("Value of p: ", *p)
	fmt.Println("Address of p: ", p)

	*p = 100 // assign value to allocated memory

	// After assigning value
	fmt.Println("Value of p: ", *p)
	fmt.Println("Address of p: ", p)

	newArray := new([5]int) // allocate memory for an array of 5 integers
	fmt.Println("Lenght: ", len(newArray), " Capacity: ", cap(newArray))

	newArray[0] = 10 // assign a value to first element
	newArray[4] = 20 // assign value to last element
	fmt.Println("New array: ", *newArray)
	fmt.Println("New array: ", newArray)


	newVertex := new([5]Vertex) // allocate memory for a Vertex
	fmt.Println("Lenght: ", len(newVertex), " Capacity: ", cap(newVertex))
	// assign a value to first element
	newVertex[0].X = 1
	newVertex[0].Y = 10
	// assign value to last element
	newVertex[4].X = 5
	newVertex[4].Y = 10
	fmt.Println("New Vertex: ", newVertex)

	newStruct := new(Vertex)
	// accessing the element
	newStruct.X = 10
	newStruct.Y = 20
	fmt.Println("New Struct: ", newStruct)
}
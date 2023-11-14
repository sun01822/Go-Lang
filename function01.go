package main

import (
	"fmt"
)

// implement the function sum(x,y) here
func sum(x int, y int) int {
	return x+y
}

func main() {
	x,y := 10, 20
	result := sum(x,y)
	fmt.Println("The sum of ", x, " and ", y, " is ", result)
}
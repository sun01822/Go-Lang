package method

import (
	"fmt"
)

type MyString string

type Rectangular struct {
	Width float64
	Height float64
}

func EmptyInterface(I interface{}){
	fmt.Printf("Value given to empty interface function is of type '%T' with value %v\n", I, I)
}
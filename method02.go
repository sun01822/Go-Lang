package main

import (
	"fmt"
)

type myNumber int   // non struct data type


func (num myNumber) square() myNumber{
	if num == 0{
		return 1
	}
	return num * num
}


func main() {
	num := myNumber(25)
	sq:= num.square()
	fmt.Printf("The square of %d is %d\n", num, sq)

}
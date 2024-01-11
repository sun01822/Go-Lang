package main

import (
	"fmt"
	"strconv"
)

func DivisibleByThree(input int) string{
	isDivisibleByThree := (input % 3) == 0
	if isDivisibleByThree{
		return "Hurray!"
	}
	return strconv.Itoa(input)
}

func main(){
	fmt.Println(DivisibleByThree(3))
}
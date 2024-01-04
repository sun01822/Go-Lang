package main

import (
	"fmt"
)

type MyString string 

type Rectangular struct{
	width, height float64
}

// emptyInterface function accepts an empty interface as a parameter
func emptyInterface(i interface{}){
	fmt.Printf("Value given to emptyInterface function is of type '%T' with value %v\n", i, i)
}

func main(){
	var str MyString = "Hello World"
	emptyInterface(str)
	var rectangular = Rectangular{10.20, 20.3}
	emptyInterface(rectangular)
	floatNum := 20.5
	var integerVal = int64(floatNum)
	emptyInterface(integerVal)
}
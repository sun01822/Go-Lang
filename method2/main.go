package main

import (
	"fmt"
)

type myNumber int 

func (num myNumber) square() myNumber{
	return num * num
}

func main(){
	var num1 myNumber = 5
	num2 := myNumber(0)
	sq := num2.square()
	fmt.Printf("The square of %d is %d\n",num1,num1.square())
	fmt.Printf("The square of %d is %d\n",num2,sq)	
}
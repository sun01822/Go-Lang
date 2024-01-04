package main

import (
	"fmt"
)

type myNumber int 

func (num myNumber) square() myNumber{
	return num * num
}

func main(){
	num := myNumber(10)
	sq := num.square()
	fmt.Printf("The square of %d is %d\n",num,sq)	
}
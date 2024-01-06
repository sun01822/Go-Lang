package main

import (
	"fmt"
)

func main(){
	var ch chan int
	if ch == nil{
		fmt.Println("channel is nil, going to define it")
		ch = make(chan int)
		fmt.Printf("Type of ch is %T", ch)
	}
}

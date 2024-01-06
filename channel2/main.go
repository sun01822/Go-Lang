package main

import (
	"fmt"
)

func hello(done chan bool){
	fmt.Println("Hello world goroutine")
	done <- false  // signal that the work is done
}

func main(){
	done := make(chan bool)
	go hello(done)
	if <-done{ // wait for goroutine to finish
		fmt.Println("Channel received true")
	}else{
		fmt.Println("Channel received false")
	}
	
	fmt.Println("main function")
}

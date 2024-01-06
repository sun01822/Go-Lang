package main

import (
	"fmt"
	"time"
)

func hello(done chan bool){
	fmt.Println("hello Go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello Go routine awake and going to write to done")
	done <- true // signal that the work is done
}

func main(){
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<- done // waiting for the signal from the go routine
	fmt.Println("Main received data")
	fmt.Println("Main executed successfully")
}

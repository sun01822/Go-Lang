package main

import (
	"fmt"
	"time"
)


func worker(cancel <-chan struct{}){
	fmt.Println("worker started")
	defer fmt.Println("worker ended")
	for{
		select{
			case <-cancel:
				fmt.Println("worker cancelled")	
				return
			default:
				time.Sleep(1*time.Second)
				fmt.Println("working...")
		}
	}
}


func main(){
	cancel := make(chan struct{})
	go worker(cancel)
	time.Sleep(5*time.Second)
	fmt.Println("main: cancelling worker goroutine")
	close(cancel)
	time.Sleep(2*time.Second)
	fmt.Println("main: done")
}
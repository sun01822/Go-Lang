package main


import (
	"fmt"
	"time"
)

func worker(cancel <-chan struct{}){
	fmt.Println("worker started")
	defer fmt.Println("worker stopped")
	for{
		select{
		default:
			// Do some work here
			time.Sleep(time.Second)
			fmt.Println("Working...")
		case <-cancel:
			fmt.Println("Worker canceled")
			return
		}
	}
}


func main() {
	cancel := make(chan struct{})
	go worker(cancel)
	time.Sleep(5 * time.Second)
	fmt.Println("Cancelling worker goroutine...")
	close(cancel)
	time.Sleep(time.Second * 2)
	fmt.Println("Main goroutine stopped")
}
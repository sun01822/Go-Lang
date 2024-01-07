package main

import (
	"fmt"
	"sync"
	"time"
)


func routine(waitgroup *sync.WaitGroup, number int){
	defer waitgroup.Done()
	fmt.Println("Start routine ", number)
	time.Sleep(time.Second)
	fmt.Println("End routine ", number)
}


func main() {
	fmt.Println("Starting main Goroutine")
	var waitgroup sync.WaitGroup
	for i := 0; i < 5; i++ {
		waitgroup.Add(1)
		go routine(&waitgroup, i)
	}
	waitgroup.Wait()
	fmt.Println("Finishing main Goroutine")
}
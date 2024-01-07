package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	fmt.Println("Starting main Goroutine")
	var waitgroup sync.WaitGroup
	waitgroup.Add(5)
	for i := 0; i < 5; i++ {
		go func(waitgroup *sync.WaitGroup, number int){
			defer waitgroup.Done()
			fmt.Println("Start routine ", number)
			time.Sleep(time.Second)
			fmt.Println("End routine ", number)
		}( &waitgroup, i)
	}
	waitgroup.Wait()
	fmt.Println("Finishing main Goroutine")
}
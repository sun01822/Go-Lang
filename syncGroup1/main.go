package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex

func main(){
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		time.Sleep(time.Duration(1) * time.Second)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func increment(wg *sync.WaitGroup){
	fmt.Println("Incrementing...")
	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()
	defer fmt.Println("Incremented")
	counter++
	fmt.Println("Incremented: ", counter)
}

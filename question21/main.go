package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println(value)
		}
	}()
	wg.Wait()
}

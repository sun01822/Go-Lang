package main

import (
	"fmt"
	"sync"
)

func test1(ch chan<- int, wg *sync.WaitGroup){
	defer wg.Done()
	arr := []int{1,2,3,4,5, 6, 7, 8, 9, 10}
	for _, value := range arr{
		ch <- value
	}
	close(ch)
}

func test2(ch <-chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for value := range ch{
		fmt.Println(value)
	}
}

func main(){
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go test1(ch, &wg)
	go test2(ch, &wg)
	wg.Wait()
	fmt.Println("Done")
}
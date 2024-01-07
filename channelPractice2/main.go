package main

import (
	"fmt"
	"sync"
)

func getValueAndPrintResult(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	value := <-ch
	fmt.Println("The square of the number ", value, " is: ", value*value)
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	wg.Add(1)
	go getValueAndPrintResult(ch, &wg)
	ch <- input
	wg.Wait()
	fmt.Println("Done")
}
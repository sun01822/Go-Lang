package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, val := range s {
		sum += val
	}
	c <- sum
}

func main() {
	fmt.Println("Hello, world!")
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	c := make(chan int)
	go sum(s1[:5], c)
	go sum(s1[5:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}



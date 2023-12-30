package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func divideSum(s int, c2 chan bool) {
	if s%2 == 0 {
		c2 <- true
	} else {
		c2 <- false
	}
	// send sum to c
}

func main() {
	s1 := []int{7, 2, 8, -9, 4, 0}
	s2 := []int{8, 9, 1, 3, 5}
	ch := make(chan int)
	ch2 := make(chan bool)
	go sum(s1, ch)
	go sum(s2, ch)
	x, y := <-ch, <-ch // receive from c
	sumOfSlices := x + y
	go divideSum(sumOfSlices, ch2)
	ans := <-ch2
	if ans {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

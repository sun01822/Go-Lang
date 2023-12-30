package main

import (
	"fmt"
	"time"
)

func PrintDigits() {
	for i := 1; i <= 5; i++ {
		fmt.Print(i)
	}
}

func main() {
	PrintDigits()
	go PrintDigits()
	time.Sleep(2 * time.Second)
}

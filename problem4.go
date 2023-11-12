package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Enter a positive number: ")
	fmt.Scanf("%d", &n)
	sum := 0
	
	// write the loop to calculate the summation here
	for i := 1; i <= n; i++ {
		if(i%3 == 0 || i%5 == 0) {
			sum += i
		}
	}
	fmt.Println("The sum of all integers from 1 to ", n, " that are divisible by either 3 or 5 is ", sum)
}
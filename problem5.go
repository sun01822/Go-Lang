package main

import (
	"fmt"
)

func main() {
	nums := []int{}
	var n int
	fmt.Print("Enter the number of elements in the slice: ")
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		var num int
		fmt.Print("Enter an integer: ")
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("The sum of all elements in the slice is: ", sum)
}
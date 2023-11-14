package main

import (
	"fmt"
)

// implemetation the function sumAll(arg...  int) here
func sumAll(arg... int) int {
	sum := 0
	for _, num := range arg {
		sum+=num
	}
	return sum
}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 2, 3, 4}
	arr3 := []int{1, 2, 3, 4, 5}

	result1 := sumAll(arr1...);
	result2 := sumAll(arr2...);
	result3 := sumAll(arr3...);

	fmt.Println("Sum of first array ", result1)
	fmt.Println("Sum of second array ", result2)
	fmt.Println("Sum of third array ", result3)
}
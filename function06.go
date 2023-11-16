package main

import (
	"fmt"
)


func average(numbers []int) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += float64(number)
	}
	return sum / float64(len(numbers))
}

func main(){
	numbers := []int{1, 2, 3, 4, 5}
	
	avg := average(numbers)
	fmt.Printf("The average of %v is %.2f\n", numbers, avg)
}
package main

import (
	"fmt"
)

func average(arr []int) float64{
	sum := 0
	for _, value := range arr{
		sum+=value
	}
	return float64(sum)/float64(len(arr))
}

func main(){
	var n int 
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter the elements: ")
	for i:=0; i<n; i++{
		fmt.Scan(&arr[i])
	}
	avg := average(arr)
	fmt.Println("The average of the slice is: ", avg)
}

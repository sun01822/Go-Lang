package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	arr[4] = 100

	fmt.Println("set: ", arr)
	fmt.Println("get: ", arr[4])
}
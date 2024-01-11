package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from runtime error:", r)
		}
	}()

	fmt.Println("Starting program")
	var myArray [3]int
	pos := 3
	myArray[pos] = 1
	fmt.Println("Program finished.")
}

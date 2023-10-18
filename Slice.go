package main

import (
	"fmt"
)

func main() {
	myslice := []int{1, 2, 3, 4, 5}
	fmt.Println(myslice)

	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(myslice2)
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))

	// Create a Slice from an array
	var myarray = []int{1, 2, 3, 4, 5}
	myslice3 := myarray[0:3]
	fmt.Println(myslice3)
	fmt.Printf("Length = %d\n", len(myslice3))
	fmt.Printf("Capacity = %d\n", cap(myslice3))

	// Create a Slice with the make() Function
	myslice4 := make([]int, 3, 10)
	fmt.Println(myslice4)
	fmt.Printf("Length = %d\n", len(myslice4))
	fmt.Printf("Capacity = %d\n", cap(myslice4))

	// with omitted Capacity
	myslice5 := make([]int, 5)
	fmt.Println(myslice5)
	fmt.Printf("Length = %d\n", len(myslice5))
	fmt.Printf("Capacity = %d\n", cap(myslice5))

	// access elements of a Slice
	fmt.Println(myslice3[0])
	fmt.Println(myslice3[1])

	myslice10 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice10 = %v\n", myslice10)
	fmt.Printf("Length = %d\n", len(myslice10))
	fmt.Printf("Capacity = %d\n", cap(myslice10))

	myslice10 = append(myslice10, 20, 21)
	fmt.Printf("\nmyslice10 = %v\n", myslice10)
	fmt.Printf("Length = %d\n", len(myslice10))
	fmt.Printf("Capacity = %d\n", cap(myslice10))

	//Append one slice to another slice
	myslice6 := []int{1, 2, 3}
	myslice7 := []int{4, 5, 6}

	myslice8 := append(myslice6, myslice7...)

	fmt.Printf("\nmyslice8 = %v\n", myslice8)
	fmt.Printf("Length = %d\n", len(myslice8))
	fmt.Printf("Capacity = %d\n", len(myslice8))

	// Change the lenght of a slice
	array1 := [6]int{9, 10, 11, 12, 13, 14} // an array
	myslice11 := array1[1:5]                // Slice array

	fmt.Printf("\nmyslice11 = %v\n", myslice11)
	fmt.Printf("Length = %d\n", len(myslice11))
	fmt.Printf("Capacity = %d\n", cap(myslice11))

	myslice11 = array1[1:3] // Change length by re-slicing the array
	fmt.Printf("\nmyslice11 = %v\n", myslice11)
	fmt.Printf("Length = %d\n", len(myslice11))
	fmt.Printf("Capacity = %d\n", cap(myslice11))

	// Copy Function
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// Original Slice
	fmt.Printf("\nmyslice11 = %v\n", numbers)
	fmt.Printf("Length = %d\n", len(numbers))
	fmt.Printf("Capacity = %d\n", cap(numbers))

	// Create copy with only nedded numbers
	needNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(needNumbers))
	copy(numbersCopy, needNumbers)

	fmt.Printf("\nnumbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}

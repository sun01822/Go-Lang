package main

import ("fmt")

func main() {
	
	var arr1 = [3]int{1, 2, 3}  // Normal initalization
	arr2 := [5]int{1, 2, 3, 4, 5} // using := operator

	fmt.Println(arr1)
	fmt.Println(arr2)


	var arr3 = [...]int{1,2,3}
	arr4 := [...]int{1,2,3,4,5,6}

	fmt.Println(arr3)
	fmt.Println(arr4)


	// String array
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	
	// Acess elements of an array
	prices := [3]int{10,20,30}

	fmt.Println(prices[0])
	fmt.Println(prices[1])

	// Change elements of an array
	prices[1] = 1000
	
	fmt.Println(prices[1])

	
	arrr1 := [5]int{}  			// not initalization
	arrr2 := [5]int{1,2} 		// partial initialized
	arrr3 := [5]int{1,2,3,4,5}   // fully initialized


	fmt.Println(arrr1)
	fmt.Println(arrr2)
	fmt.Println(arrr3)


	// initialize only specific elements
	arr5 := [5]int{1:10, 2:40}

	fmt.Println(arr5)

	arrr4 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arrr5 := [...]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14}

	fmt.Println(arrr4)
	fmt.Println(arrr5)


	// find the length of an array
	fmt.Println(len(arrr4))
	fmt.Println(len(arrr5))

}
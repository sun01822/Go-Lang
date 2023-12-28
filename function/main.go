package main

import (
	"fmt"
)

func vals()(int, string) {
	return 3, "Another variable"
}

func arrayFunction(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

// varadic function 
func variadicFunction(s string, nums ...int) {
	sum := 0
	diff := 0
	if s == "sum" {
		fmt.Print("Sum of the numbers is: ")
		for _, num := range nums {
			sum+=num
		}
		fmt.Println(sum)
	}else{
		fmt.Print("Difference of the numbers is: ")
		for _, num := range nums {
			diff-=num
		}
		fmt.Println(diff)
	}
}


// return parameter name
func calculationFunc(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return
}


func main() {
	val, str := vals()
	fmt.Println(val)
	fmt.Println(str)

	_, str2 := vals()  // escape the first variable
	fmt.Println(str2)

	arr := []int{1, 2, 3}
	arrayFunction(arr)
	fmt.Println(arr)
	variadicFunction("sum", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	val1, val2 := calculationFunc(10, 5)
	fmt.Println(val1)
	fmt.Println(val2)
}
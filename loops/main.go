package main

import (
	"fmt"
)

func main() {
	for{
		fmt.Println("Hello")
		break
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	num:= make([]int, 0, 10);
	sum:=0
	num = append(num, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for _, n := range num {
		sum+=n
	}
	fmt.Println(sum)

	// implementing a map[type]type
	m := map[string]string{
		"a": "apple", 
		"b": "banana",
	}
	for k, v := range m {
		fmt.Println(k," -> ",v)
	}
}
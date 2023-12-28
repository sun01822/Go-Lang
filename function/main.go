package main

import (
	"fmt"
)

func vals()(int, string) {
	return 3, "Another variable"
}

func main() {
	val, str := vals()
	fmt.Println(val)
	fmt.Println(str)

	_, str2 := vals()  // escape the first variable
	fmt.Println(str2)
}
package main


import (
	"fmt"
)

func main() {
	var_1, var_2 := 1, "hi"
	fmt.Println(var_1, var_2)
	var_3, var_2 := 2, "hello"
	fmt.Println(var_3, var_2)
	// same scope, the variable is reassigned
	if var_4, var_2 := 3, "hey"; var_4 > 2 {
		// variable is declared again in the scope of the if condition
		fmt.Println(var_4, var_2)
	}
	fmt.Println(var_2)
	// fmt.Println(var_4) // error: undefined: var_4
}
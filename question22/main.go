package main

import (
	"fmt"
)

func main() {
	add := func(a *int, b int) {
		*a = *a + b
		return
	}
	a := 5
	add(&a, 2)
	fmt.Println(a)
}

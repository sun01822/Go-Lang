package main

import (
	"fmt"
)

func main() {
	var map1 map[string]int
	if map1 == nil {
		fmt.Println("map1 is nil")
	} else {
		fmt.Println("map1 is not nil")
	}
	map2 := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Fish",
	}
	fmt.Println(map2)
	fmt.Println(map2[91])
	fmt.Println(len(map2))
	delete(map2, 93)
	fmt.Println(map2)
}
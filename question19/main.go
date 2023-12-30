package main

import (
	"fmt"
)

func main() {
	myMap := map[string]interface{}{
		"a": 42,
		"b": "hello",
		"c": []int{1, 2, 3},
	}
	sum := myMap["a"].(int) + len(myMap["b"].(string)) + len(myMap["c"].([]int))
	fmt.Println(sum)
}

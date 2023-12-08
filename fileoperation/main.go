package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var boolean = true
	bolB, _ := json.Marshal(boolean)
	for i:=range bolB{
		fmt.Print(string(bolB[i]))
	}
	fmt.Println()
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.1416)
	fmt.Println(string(fltB[3]))

	// Bhaiya der jiggas korte hobe
	str := "Sun"
	strB, _ := json.Marshal(str) // Output: S
	// strB, _ := json.Marshal("Sun")  // Output: "
	fmt.Println(string(strB[1]))

	slcD := []string{"apple", "peach", "pear"}
	fmt.Println("before: ", slcD)
	slcB, err := json.Marshal(slcD)
	if err != nil {
		fmt.Println("after: ", slcB) 
	}


	// map to JSON 
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) 


	// struct to JSON 
	type response1 struct {
		Page int
		Fruits []string
	}

	res1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	type response2 struct{
		Page int 	`json:"page-number"`
		Fruits []string `json:"fruits-List"`
	}

	res2D := &response2{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(
		`
		{
			"num": 6.13,
			"strs": ["a", "b", "c"]
		}
		`,
	)

	var v map[string]interface{}
	if err := json.Unmarshal(byt, &v); err != nil {
		panic(err)
	}
	fmt.Println(v)
	
}
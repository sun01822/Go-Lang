package main

import (
	"fmt"
	"encoding/json"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int `json:"page-number"`
	Fruits []string `json:"fruits-list"`
}

func main(){
	// Boolean values encode as JSON booleans.
	var boolean = true
	b, err := json.Marshal(boolean)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println(string(b))
	}

	// integer values encode as JSON numbers.
	var integer = 1
	i, err := json.Marshal(integer)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println(string(i))
	}

	// floating point encode as JSON numbers.
	var float = 1.23
	f, err := json.Marshal(float)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println(string(f))
	}

	// string values encode as JSON strings
	var str = "Hello World"
	s, err := json.Marshal(str)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println(str)
		fmt.Println(string(s[0]))
		fmt.Println(string(s))
	}

	// direct string encode as JSON strings
	d, err := json.Marshal("Hello World")
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println(string(d[0]))
		fmt.Println(string(d))
	}


	// slice encode as JSON array
	sliceFruit := []string{"apple", "banana", "mango"}
	fmt.Println("Before: ", sliceFruit)
	slice, err := json.Marshal(sliceFruit)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("After: ",string(slice))
	}

	// map encode as JSON object
	mapFruit := map[string]int{"apple": 5, "banana": 2, "mango": 1}
	fmt.Println("Before: ", mapFruit)
	mapF, err := json.Marshal(mapFruit)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("After: ",string(mapF))
	}

	res1D := &Response1{
		Page: 1,
		Fruits: []string{"apple", "banana", "mango"},
	}
	fmt.Println("Before: ", res1D)
	res1B, err := json.Marshal(res1D)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("After: ", string(res1B))
	}

	res2D := &Response2{
		Page: 1,
		Fruits: []string{"apple", "banana", "mango"},
	}
	fmt.Println("Before: ", res2D)
	res2B, err := json.Marshal(res2D)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("After: ", string(res2B))
	}
}
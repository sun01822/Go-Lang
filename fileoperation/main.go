package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var boolean = true
	bolB, _ := json.Marshal(boolean)
	for i := range bolB {
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
		Page   int
		Fruits []string
	}

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	type response2 struct {
		Page   int      `json:"page-number"`
		Fruits []string `json:"fruits-List"`
	}

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//  Unmarshal JSON response
	fmt.Println("\nUse of Unmarshal:")
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

	type response3 struct {
		Page   int      `json:"number"`
		Fruits []string `json:"fruits-List"`
	}

	str2 := `
	{
    	"number": 1,
    	"fruits-List": ["apple", "peach"]
	}`

	res := response3{}
	json.Unmarshal([]byte(str2), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// JSON validation
	fmt.Println("\nJSON Validation Checking")
	type response4 struct {
		Page   int      `json:"page-number"`
		Fruits []string `json:"fruits-List"`
	}

	str3 := `
	{
    	"number": 1, 
		"fruits-List": ["apple", "peach"]
	}`
	/**

		str3 := `
		{
	    	"number": 1 "fruits-List": ["apple", "peach"]
		}`
	*/

	if json.Valid([]byte(str3)) {
		res := response4{}
		json.Unmarshal([]byte(str3), &res)
		fmt.Println(res)
		fmt.Println(res.Fruits[0])
	} else {
		fmt.Println("Wrong JSON response")
	}

	// JSON encoding
	/*
		first we need to import the packages "encoding/json" and "os"
	*/
	fmt.Println("\nJSON encoding:")
	person := Person{Name: "Sun", Age: 25}
	file, err := os.Create("person.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON file created")
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err = encoder.Encode(person); err != nil {
		panic(err)
	}
	fmt.Println("JSON encoding done")


	// Reading JSON data using NewEncoder()
	fmt.Println("\nJSON decoding:")
	if readFile, err := os.Open("person.json"); err != nil {
		panic(err)
	}else{
		fmt.Println("JSON file opened successfully")
		defer readFile.Close()
		decoder := json.NewDecoder(readFile)
		var person Person
		if err = decoder.Decode(&person); err != nil {
			panic(err)
		}
		fmt.Println("JSON decoding done")
		fmt.Println(person)
		fmt.Println("Decoded person Name: ", person.Name)
		fmt.Println("Decoded person Age: ", person.Age)
	}
}

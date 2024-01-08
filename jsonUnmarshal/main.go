package main

import (
	"fmt"
	"encoding/json"
)

type Response struct{
	Page int `json:"page-number"`
	Fruits []string `json:"fruits-List"`
}

func main(){
	
	// using this could be runtime error
	byt := []byte(
		`{
			"num":6.13,
			"strs":["a","b"]
			}`,
		)
	
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil{
		panic(err)
	}else{
		fmt.Println(dat)
	}

	// to remove runtime error use this
	str := `{
		"page-number":1,
		"fruits-List":["apple","peach"]
	}`

	res := Response{}
	err :=  json.Unmarshal([]byte(str), &res)
	if err != nil{
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	
	// checking valid json or not
	str2 := `{
		"page-number":10 "fruits-List":["apple","peach", "mango"]
	}`
	// missing a comma between elements
	if json.Valid([]byte(str2)){
		fmt.Println("Valid JSON Response")
		res2 := Response{}
		err := json.Unmarshal([]byte(str2), &res2)
		if err != nil{
			panic(err)
		}
		fmt.Println(res2)
		fmt.Println(res2.Fruits[2])
	}else{
		fmt.Println("Wrong JSON Response")
	}
}

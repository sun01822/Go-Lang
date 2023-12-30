package main

import (
	"fmt"
	"encoding/json"
)

type response2 struct {
	Page   int      `json:"page_number"`
	Fruits []string `json:"fruits-list"`
}

func main() {
    var	new string = "new"
	jsonStr := `
		{
			"page-number": 1,
			"fruits-list": ["apple", "peach"]
		}`

	res := response2{}
	json.Unmarshal([]byte(jsonStr), &res)
	fmt.Println(res)
}

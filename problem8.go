package main

import (
	"fmt"
)

func main() {
	var day string = "Saturday"
	switch day{
		case "Saturday": 
		fmt.Println(day, " is workday")
		case "Sunday": 
		fmt.Println(day, " is workday")
		case "Monday": 
		fmt.Println(day, " is workday")
		case "Tuesday":  
		fmt.Println(day, " is workday")
		case "Wednesday":  
		fmt.Println(day, " is workday")
		case "Thursday":  
		fmt.Println(day, " is workday")
		case "Friday":  
		fmt.Println(day, " is holiday")
	default:
		fmt.Println("Unknown day")
	}
}
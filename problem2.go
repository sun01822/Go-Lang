package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Printf("Enter a year: ")
	fmt.Scanf("%d",&year)

	// Check if the year is a leap year or not
	if year % 4 == 0 {
		if(year % 100 == 0){
			if(year % 400 == 0){
				fmt.Printf("%d is a leap year\n", year)
			}else{
				fmt.Printf("%d is not a leap year\n", year)
			}
		}else{
			fmt.Printf("%d is a leap year\n", year)
		}
	}else{
		fmt.Printf("%d is not a leap year\n", year)
	}
}
package main

import "fmt"

func main() {
	var day int
	fmt.Printf("Enter a day: ")
	fmt.Scanf("%d", &day)

	dayName := ""

	switch day {
		case 1: dayName = "Monday"
		case 2: dayName = "Tuesday"
		case 3: dayName = "Wednesday"
		case 4: dayName = "Thursday"
		case 5: dayName = "Friday"
		case 6: dayName = "Saturday"
		case 7: dayName = "Sunday"
		default: dayName = "Nothing"
	}
	fmt.Printf("Name of the day is %s", dayName)
}
package main

import (
	"fmt"
	"math"
)


// implement check(n) fmt function here
func check(n int) (float64, string){
	if n < 0{
		return 0 ,"negative"
	}else{
		return math.Sqrt(float64(n)), "positive"
	}
}


func main(){
	var n int
	fmt.Printf("Enter the number: ")
	fmt.Scan(&n)
	result, message := check(n)
	if message == "negative"{
		fmt.Println("This is a negative number, So square root of ", n, "is impossible")
	}else{
		fmt.Println("Square root of ", n, "is ", result)
	}
}
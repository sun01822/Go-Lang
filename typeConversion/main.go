package main

import (
	"fmt"
)

func main(){
	var iNum int = 10
	var fNum float32 = float32(iNum)
	fmt.Printf("value = %v | type - %T\n", fNum, fNum)
	var fNum2 = 5.65
	iNum1 := int(fNum2)
	fmt.Printf("value = %v | type - %T\n", iNum1, iNum1)

	var num1 float64 = 10.5
	var num2 int = 5
	result := float64(num1) + float64(num2)
	fmt.Printf("value = %v | type - %T\n", result, result)


	// formating numbers
	var d1 int = 10
	fmt.Printf("in binary=%b\n", d1)
	fmt.Printf("in decimal=%d\n", d1)
	fmt.Printf("in hex=%x\n", d1)
	fmt.Printf("in octal=%o\n", d1)

}

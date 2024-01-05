package main

import (
	"fmt"
	basic "practiceProblem6/calculator/basic"
	scientific "practiceProblem6/calculator/scientific"
)

func main(){
	var a float64 = 10
	var b float64 = 5
	var c float64 = 2
	var d float64 = 3


	basicCalc := basic.BasicCalculator{FirstNumber: a, SecondNumber: b}
	scientificCalc := scientific.ScientificCalculator{FirstNumber: c, SecondNumber: d}


	fmt.Println("Basic Calculator")
	fmt.Println("Add: ", basicCalc.Add())
	fmt.Println("Subtract: ", basicCalc.Subtract())
	fmt.Println("Multiply: ", basicCalc.Multiply())
	fmt.Println("Divide: ", basicCalc.Divide())

	fmt.Println("Scientific Calculator")
	fmt.Println("Add: ", scientificCalc.Add())
	fmt.Println("Subtract: ", scientificCalc.Subtract())
	fmt.Println("Multiply: ", scientificCalc.Multiply())
	fmt.Println("Divide: ", scientificCalc.Divide())
}

package main

import ("fmt")

func main() {
    var a = 15 + 15
	fmt.Println(a)


	// & operator
	var x = 5
	fmt.Printf("x is %b\n", x)  // 101
	fmt.Printf("3 is %03b\n", 3)  // 011

	
	x&=3
	fmt.Printf("x is now %03b\n", x)  // 001


	//  | operator
	var y = 5
	fmt.Printf("y is %b\n", y)  // 101
	fmt.Printf("3 is %03b\n", 3)  // 011

	
	y|=3
	fmt.Printf("y is now %03b\n", y)  // 111


	// ^ operator
	var z = 5

	fmt.Printf("z is %b\n", z)  // 101
	fmt.Printf("3 is %03b\n", 3)  // 011

	
	z^=3
	fmt.Printf("z is now %03b\n", z)  // 110



}
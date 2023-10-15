package main

import "fmt"

func main() {
	c1 := complex(2,3)     // complex initializer syntax a+ib
	c2 := 4 + 5i           // additional just like other variables
	c3 := c1 + c2          // prints "Add: (6+8i)"
	fmt.Println("Add: ", c3)

	re := real(c3)   // get real part
	im := imag(c3)   // get imaginary part

	fmt.Println(re, im)  // print 6 8
}
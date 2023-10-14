package main

import "fmt"

func main() {
    // signed integers
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	// unsigned integers
	var e uint8 = 255
    var f uint16 = 65535
    var g uint32 = 4294967295
    var h uint64 = 18446744073709551615


	// Print the values
	fmt.Println("Signed Integers: ")
	fmt.Printf("int8: %d\n", a)
	fmt.Printf("int16: %d\n", b)
	fmt.Printf("int32: %d\n", c)
	fmt.Printf("int64: %d\n", d)
	fmt.Println("Unsigned Integers: ")
	fmt.Println("unit8: %d\n",e)
	fmt.Println("uint16: %d\n",f)
	fmt.Println("uint32: %d\n",g)
	fmt.Println("uint64: %d\n",h)   

}
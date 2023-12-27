package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
	fmt.Println("zeroval in function:", ival)
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i:= 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i) // i এর অ্যাড্রেস প্রিন্ট হবে
}
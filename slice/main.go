package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println("arr:", arr)
	fmt.Println("slice:", slice)

	slice2 := make([]string, 5)
	slice2[0] = "a"
	slice2[1] = "b"
	slice2[2] = "c"
	slice2[3] = "d"
	slice2[4] = "e"
	fmt.Println("set: ", slice2)
	fmt.Println("get: ", slice2[2])
	fmt.Println("len: ", len(slice2))
	slice2 = append(slice2, "f")
	slice2 = append(slice2, "g", "h")
	fmt.Println("apd: ", slice2)
	fmt.Println("len: ", len(slice2))


	s:=make([]string,3)
	s[0]="a"
	s[1]="b"
	s[2]="c"
	s = append(s,"d")
	s = append(s,"e","f")
	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("cpy:",c)

	
}
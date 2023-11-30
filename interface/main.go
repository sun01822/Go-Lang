package main

import (
	"fmt"
	"interface/method"
)

func main() {
	mystring := method.MyString("My name is MD. Shariar Hossain Sun")
	rectangular := method.Rectangular{Width:5.2, Height: 7}

	method.EmptyInterface(mystring)
	method.EmptyInterface(rectangular)

	var c method.Cylinder
	c.Radius = 3.2
	c.Height = 4.3
	fmt.Println(c.Area())
	fmt.Println(c.Volume())

	var shape method.Shape = method.Cylinder{3.2, 4.3}
	cylinder := shape.(method.Cylinder)
	fmt.Println("Area of shape of interface type Shape is ", cylinder.Area())
    fmt.Println("Volume of object of interface type Object is ", cylinder.Volume())

}
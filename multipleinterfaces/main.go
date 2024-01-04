package main

import (
	"fmt"
)

type Shape interface{
	Area() float64
}

type Object interface{
	Volume() float64
}

type Cylinder struct{
	radius float64
	height float64
}

func (c Cylinder) Area() float64{
	return 2*3.14*c.radius*c.height
}

func (c Cylinder) Volume() float64{
	return 3.14*c.radius*c.radius*c.height
}

func main(){
	var shape Shape = Cylinder{10,10}
	cylinder := shape.(Cylinder)

	var object Object = Cylinder{10,10}
	cylinder2 := object.(Cylinder)

	fmt.Println("Area of cylinder is", cylinder.Area())
	fmt.Println("Volume of cylinder is", cylinder.Volume())
	fmt.Println("Area of cylinder is", cylinder2.Area())
	fmt.Println("Volume of cylinder is", cylinder2.Volume())

	var cylinder3 Cylinder = shape.(Cylinder)
	fmt.Println("Area of cylinder is", cylinder3.Area())
	fmt.Println("Volume of cylinder is", cylinder3.Volume())

}


package formula

import "fmt"

type Geometry interface {
	Area() float64
	Perimeter() float64
}

func Calculate(g Geometry) {
	fmt.Println("Info of geometry: ", g)
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}
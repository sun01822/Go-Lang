package main

import (
	p "practiceProblem3/formula"
)

type rectangle struct {
	length float64
	width float64
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}

func (r rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main(){
	r := rectangle{length: 10, width: 5}
	c := circle{radius: 12}
	p.Calculate(r)
	p.Calculate(c)
}
package method

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Cylinder struct{
	Radius float64
	Height float64
}

func (c Cylinder) Area() float64{
	return 2 * math.Pi * c.Radius * (c.Radius + c.Height)
}

func (c Cylinder) Volume() float64{
	return math.Pi * c.Radius * c.Radius * c.Height
}
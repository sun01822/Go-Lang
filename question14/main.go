package main

import (
	"fmt"
)

type Mover interface {
	Move() string
}
type Boat struct{}

func (b Boat) Move() string {
	return "Sailing"
}

type Car struct{}

func (c Car) Move() string {
	return "Driving"
}

func main() {
 movers := []Mover{Car{}, Boat{}} // A slice of Mover
 for _, m := range movers {
	 fmt.Println(m.Move())
 }
}

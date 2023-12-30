package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func modifyPoint(p Point, newX, newY int) Point {
	p.X = newX
	p.Y = newY
	return p
}

func main() {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 3, Y: 4}
	p3 := modifyPoint(p1, 10, 20)
	fmt.Printf("p1: %v, p2: %v, p3: %v\n", p1, p2, p3)
}

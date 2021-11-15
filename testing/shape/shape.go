package main

import (
	"math"
)

// func main() {

// }

type Rectangle struct {
	length, breadth float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}



func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}



func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

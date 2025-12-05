package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func PrintShapeMethod(s Shape) {
	fmt.Println(s.Area())
}
func main() {

	r1 := Rectangle{
		Width:  10,
		Height: 20,
	}
	c1 := Circle{
		Radius: 10,
	}
	PrintShapeMethod(r1)
	PrintShapeMethod(c1)
}

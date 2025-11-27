package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func main() {
	rectangle := Rectangle{
		Width:  10,
		Height: 5,
	}
	fmt.Println("Before ", rectangle)
	fmt.Println("Area", rectangle.Area())
	Scale(&rectangle, 2)
	fmt.Println("After ", rectangle)
}
func (rect Rectangle) Area() int {
	return rect.Height * rect.Width
}
func Scale(r *Rectangle, factor int) {
	r.Height *= factor
	r.Width *= factor
}

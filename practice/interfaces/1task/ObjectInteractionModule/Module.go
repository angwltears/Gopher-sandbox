package objectinteractionmodule

import (
	"fmt"
	"math"
)

/*
-- Разработать модуль для работы с геометрическими фигурами
-- каджая фигура умеет вычислять свою площядь и периметр
-- разные фигуры реализуются по-разному,но взаимодействие с ними должно быть унифицировано через интерфейс
--Необходимо описать фигуры и функцию, которая может принимать различные виды фигур и высчитывать их площадь и периметр
*/
type Triangle struct {
	FirstAngle  float64
	SecondAngle float64
	ThirdAngle  float64
}

func (t Triangle) Area() float64 {
	HalfPerimeter := t.Perimeter() / 2
	return math.Sqrt(HalfPerimeter * (HalfPerimeter - t.FirstAngle) * (HalfPerimeter - t.SecondAngle) * (HalfPerimeter - t.ThirdAngle))
}
func (t Triangle) Perimeter() float64 {
	return t.FirstAngle + t.SecondAngle + t.ThirdAngle

}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return c.Radius * 2 * math.Pi
}

type Square struct {
	SideLength float64
}

func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}
func (s Square) Perimeter() float64 {
	return 4 * s.SideLength
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func CalcAreaAndPerimeter(m Shape) {
	fmt.Println("Площадь фигуры: ", m.Area())
	fmt.Println("Периметр фигуры: ", m.Perimeter())
}

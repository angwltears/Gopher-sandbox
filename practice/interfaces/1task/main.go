package main

import "mod1/objectinteractionmodule"

/*
-- Разработать модуль для работы с геометрическими фигурами
-- каджая фигура умеет вычислять свою площядь и периметр
-- разные фигуры реализуются по-разному,но взаимодействие с ними должно быть унифицировано через интерфейс
--Необходимо описать фигурыи функцию, которая можетпринимать различные виды фигур и высчитывать их площядь и периметр
*/
func main() {
	R1 := objectinteractionmodule.Triangle{
		FirstAngle:  5,
		SecondAngle: 10,
		ThirdAngle:  15,
	}
	C1 := objectinteractionmodule.Circle{
		Radius: 10,
	}
	S1 := objectinteractionmodule.Square{
		SideLength: 25,
	}
	objectinteractionmodule.CalcAreaAndPerimeter(R1)
	objectinteractionmodule.CalcAreaAndPerimeter(C1)
	objectinteractionmodule.CalcAreaAndPerimeter(S1)

}

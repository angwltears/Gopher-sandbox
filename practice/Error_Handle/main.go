package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type Car struct {
	Health int
}

func (c *Car) Gas() (int, error) {
	if c.Health-10 < 0 {
		return 0, errors.New("Мы не стали газовать, так как машина бы сломалась!")
	}
	c.Health -= 10
	return rand.IntN(200), nil
}
func main() {
	Mitsubishi := Car{
		Health: 120,
	}
	for {
		speed, err := Mitsubishi.Gas()
		if err != nil {
			fmt.Println("Ошибка: ", err.Error())
			break
		}
		fmt.Println("Мы газонули!!, скорость: ", speed)
	}

}

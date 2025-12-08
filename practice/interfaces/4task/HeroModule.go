package main

import "fmt"

type Hero interface {
	Attack() int
}

func AllAttackDamage(heroes []Hero) {
	sum := 0
	for _, v := range heroes {
		sum += v.Attack()
	}
	fmt.Println("Общий урон составил: ", sum)

}

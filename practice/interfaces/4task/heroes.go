package main

import "fmt"

type Warrior struct {
	Damage int
	Weapon string
}
type Mage struct {
	Mana        int
	SpellDamage int
}
type Archer struct {
	Bow    string
	Damage int
}

func (w Warrior) Attack() int {
	fmt.Println("Атакую своим оружием - ", w.Weapon, "!")
	fmt.Println("Наношу врагу урон!", w.Damage)
	return w.Damage
}
func (m Mage) Attack() int {
	fmt.Println("Трачу ману ", m.Mana, "!")
	fmt.Println("Наношу врагу урон!", m.SpellDamage)
	return m.SpellDamage
}
func (a Archer) Attack() int {
	fmt.Println("Атакую своим оружием - ", a.Bow, "!")
	fmt.Println("Наношу врагу урон!", a.Damage)
	return a.Damage
}
